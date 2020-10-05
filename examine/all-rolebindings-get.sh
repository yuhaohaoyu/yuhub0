for i in `kubectl get rolebindings -A | tail -15 | awk '{printf "%s,%s ", $1, $2;}'`; do ns=`echo $i | cut -d "," -f 1`; nm=`echo $i | cut -d "," -f 2`; echo; echo $ns      $nm; echo "===================="; kubectl get -n $ns rolebinding $nm -o yaml ; done > /tmp/role-subjects.txt 

