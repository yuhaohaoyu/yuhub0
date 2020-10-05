 2192  for i in `kubectl api-resources --verbs=list --namespaced -o name`; do kubectl get $i -n kubevirt; done
 2193  for i in `kubectl api-resources --verbs=list --namespaced -o name`; do kubectl get $i -A; done | grep kubevirt
for i in `kubectl api-resources --verbs=list --namespaced=false -o name`; do echo; echo $i;  kubectl get $i; done 

