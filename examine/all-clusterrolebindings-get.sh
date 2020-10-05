# command lines to get all the mappings from ROLE-NAME to SUBJECT-NAME for all CLUSTER-ROLE-BINDINGS

# dump complete YAML for all CLUSTER-ROLE-BINDINGS
 for i in `kubectl get clusterrolebindings | tail -51 | cut -d " " -f 1`; do kubectl get clusterrolebinding $i -o yaml; done > all-clusterrolebindings.txt

# dump only NAMES for "metadata", "roleRef", and "subjects" for all CLUSTER-ROLE-BINDINGS
 for i in `kubectl get clusterrolebindings | tail -51 | cut -d " " -f 1`; do echo; echo $i; echo "============================="; kubectl get clusterrolebinding $i -o yaml | egrep '^  name:|^. kind:|^subjects:|^metadata:|^roleRef:'; done > all-clusterrolebindings.txt

# table the NAMES for "metadata", "roleRef", and "subjects" for all CLUSTER-ROLE-BINDINGS, field separator is "="
for i in `kubectl get clusterrolebindings | tail -51 | cut -d " " -f 1`; do echo -n $i; kubectl get clusterrolebinding $i -o yaml | egrep '^  name:|^. kind:' | sed -e "s/^-/ /g" | awk '{printf " = %s", $2;}'; echo; done > all-clusterrolebindings-table.txt
