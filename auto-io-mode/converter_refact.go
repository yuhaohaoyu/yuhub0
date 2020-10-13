	volumeIndices := { "v0" :  0, "v1" : 1 }
	volumes P:=  { "v0" :  &v0, "v1" : &v1 }


	dedicatedThreads := 0	// not reaching the main loop
				// --> interact w/ threadPoolLimit, ioThreadCount
	threadPoolLimit := 1	// not reaching the main loop

	useIOThreads := false	// reaching the main loop
	autoThreads := 0	// reaching the main loop


	Adjust_IO_Mode( vmi.Spec.Domain.Devices.Disks ,&ioModes )
	Assign_IO_ThreadId( vmi.Spec.Domain.Devices.Disks, &ioThreadIds )
	Convert_To_api_Disk( vmi.Spec.Domain.Devices.Disks, domain.Spec.Devices.Disks, ioModes, ioThreadIds )

	//  

	// prep loop: count device-side thread wishes
        for _, diskDevice := range vmi.Spec.Domain.Devices.Disks {
                dedicatedThread := false
                if diskDevice.DedicatedIOThread != nil {
                        dedicatedThread = *diskDevice.DedicatedIOThread
                }
                if dedicatedThread {
                        useIOThreads = true
                        dedicatedThreads += 1
                } else {
                        autoThreads += 1
                }
        }

	// main loop
        for i, disk := range vmi.Spec.Domain.Devices.Disks {
                newDisk := Disk{}

                err := Convert_v1_Disk_To_api_Disk(&disk, &newDisk, devicePerBus, numBlkQueues)
                volume := volumes[disk.Name]
                err = Convert_v1_Volume_To_api_Disk(volume, &newDisk, c, volumeIndices[disk.Name])
		switch isNativeIOModes[i] {
		case false: 
			newDisk.Driver.IOThread = ioThreadIds[i]
		case true: 
			// do nothing
		}

                if useIOThreads {
                        ioThreadId := defaultIOThread
                        dedicatedThread := false
                        if disk.DedicatedIOThread != nil {
                                dedicatedThread = *disk.DedicatedIOThread
                        }

                        if dedicatedThread {
                                ioThreadId = currentDedicatedThread
                                currentDedicatedThread += 1
                        } else {
                                ioThreadId = currentAutoThread
                                // increment the threadId to be used next but wrap around at the thread limit
                                // the odd math here is because thread ID's start at 1, not 0
                                currentAutoThread = (currentAutoThread % uint(autoThreads)) + 1 
                        }
                        newDisk.Driver.IOThread = &ioThreadId
                }

                domain.Spec.Devices.Disks = append(domain.Spec.Devices.Disks, newDisk)
        }

