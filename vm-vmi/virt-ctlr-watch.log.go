Types:
	[root@viose2 watch]# grep '^type .* struct' *.go
	application.go:type VirtControllerApp struct {
	migration.go:type MigrationController struct {
	node.go:type NodeController struct {
	replicaset.go:type VMIReplicaSet struct {
	vm.go:type sarProxy struct {
	vm.go:type VMController struct {
	vmi.go:type syncErrorImpl struct {
	vmi.go:type VMIController struct {

Static Profiles:

	[root@viose2 watch]# wc `ls *.go | grep -v _te`
	   590   1404  20834 application.go
	   886   3195  31410 migration.go
	   314   1018   9704 node.go
	   848   3361  30402 replicaset.go
	    71    215   1596 util.go
	  1312   4917  45008 vm.go
	  1027   3944  34911 vmi.go
	  5048  18054 173865 total

Funcs:



vmi.go   =====================================

--	50	102:func NewVMIController(templateService services.TemplateService,
	4	153:func (e *syncErrorImpl) Error() string {
	2	157:func (e *syncErrorImpl) Reason() string {

	16	173:func (c *VMIController) Run(threadiness int, stopCh <-chan struct{}) {
	5	190:func (c *VMIController) runWorker() {
	18	195:func (c *VMIController) Execute() bool {

--	60	213:func (c *VMIController) execute(key string) error {

--	200	299:func (c *VMIController) updateStatus(vmi *virtv1.VirtualMachineInstance, pod *k8sv1.Pod, dataVolumes []*cdiv1.DataVolume, syncErr syncError) error {

	20	276:func conditionsEqual(a []virtv1.VirtualMachineInstanceCondition, b []virtv1.VirtualMachineInstanceCondition) bool {
	24	492:func isPodReady(pod *k8sv1.Pod) bool {
	3	516:func isPodDownOrGoingDown(pod *k8sv1.Pod) bool {
	8	520:func isComputeContainerDown(pod *k8sv1.Pod) bool {
	3	529:func podIsDown(pod *k8sv1.Pod) bool {
	6	533:func podExists(pod *k8sv1.Pod) bool {

--	50	540:func (c *VMIController) sync(vmi *virtv1.VirtualMachineInstance, pod *k8sv1.Pod, dataVolumes []*cdiv1.DataVolume) (err syncError) {

--	50	591:func (c *VMIController) handleSyncDataVolumes(vmi *virtv1.VirtualMachineInstance, dataVolumes []*cdiv1.DataVolume) (bool, syncError) {

	10	639:func (c *VMIController) addDataVolume(obj interface{}) {
	30	655:func (c *VMIController) updateDataVolume(old, cur interface{}) {
	30	687:func (c *VMIController) deleteDataVolume(obj interface{}) {

	35	716:func (c *VMIController) addPod(obj interface{}) {
	40	743:func (c *VMIController) updatePod(old, cur interface{}) {
	33	784:func (c *VMIController) deletePod(obj interface{}) {

	3	817:func (c *VMIController) addVirtualMachine(obj interface{}) {
	3	821:func (c *VMIController) deleteVirtualMachine(obj interface{}) {
	3	825:func (c *VMIController) updateVirtualMachine(old, curr interface{}) {

	10	829:func (c *VMIController) enqueueVirtualMachine(obj interface{}) {

	12	842:func (c *VMIController) resolveControllerRef(namespace string, controllerRef *v1.OwnerReference) *virtv1.VirtualMachineInstance {
	20	865:func (c *VMIController) listVMIsMatchingDataVolume(namespace string, dataVolumeName string) ([]*virtv1.VirtualMachineInstance, error) {
	20	887:func (c *VMIController) listMatchingDataVolumes(vmi *virtv1.VirtualMachineInstance) ([]*cdiv1.DataVolume, error) {
	16	907:func (c *VMIController) allPodsDeleted(vmi *virtv1.VirtualMachineInstance) (bool, error) {
	15	923:func (c *VMIController) deleteAllMatchingPods(vmi *virtv1.VirtualMachineInstance) error {
	30	954:func (c *VMIController) listPodsFromNamespace(namespace string) ([]*k8sv1.Pod, error) {
	12	967:func (c *VMIController) setActivePods(vmi *virtv1.VirtualMachineInstance) (*virtv1.VirtualMachineInstance, error) {
	24	993:func (c *VMIController) currentPod(vmi *virtv1.VirtualMachineInstance) (*k8sv1.Pod, error) {

vm.go   =====================================
	func NewVMController(vmiInformer cache.SharedIndexInformer,
	func (p *sarProxy) Create(sar *authv1.SubjectAccessReview) (*authv1.SubjectAccessReview, error) {
	func (c *VMController) Run(threadiness int, stopCh <-chan struct{}) {
	func (c *VMController) runWorker() {
	func (c *VMController) Execute() bool {
	func (c *VMController) execute(key string) error {
	func (c *VMController) handleVMRenameRequest(vm *virtv1.VirtualMachine, newName string) (bool, error) {
	func (c *VMController) listDataVolumesForVM(vm *virtv1.VirtualMachine) ([]*cdiv1.DataVolume, error) {
	func (c *VMController) orphan(cm *controller.VirtualMachineControllerRefManager, vmi *virtv1.VirtualMachineInstance) error {
	func (c *VMController) orphanDataVolumes(cm *controller.VirtualMachineControllerRefManager, dataVolumes []*cdiv1.DataVolume) error {
	func createDataVolumeManifest(dataVolumeTemplate *virtv1.DataVolumeTemplateSpec, vm *virtv1.VirtualMachine) *cdiv1.DataVolume {
	func (c *VMController) authorizeDataVolume(vm *virtv1.VirtualMachine, dataVolume *cdiv1.DataVolume) error {
	func (c *VMController) handleDataVolumes(vm *virtv1.VirtualMachine, dataVolumes []*cdiv1.DataVolume) (bool, error) {
	func (c *VMController) startStop(vm *virtv1.VirtualMachine, vmi *virtv1.VirtualMachineInstance) error {
	func (c *VMController) startVMI(vm *virtv1.VirtualMachine) error {
	func (c *VMController) stopVMI(vm *virtv1.VirtualMachine, vmi *virtv1.VirtualMachineInstance) error {
	func (c *VMController) setupVMIFromVM(vm *virtv1.VirtualMachine) *virtv1.VirtualMachineInstance {
	func setupStableFirmwareUUID(vm *virtv1.VirtualMachine, vmi *virtv1.VirtualMachineInstance) {
	func (c *VMController) filterActiveVMIs(vmis []*virtv1.VirtualMachineInstance) []*virtv1.VirtualMachineInstance {
	func (c *VMController) filterReadyVMIs(vmis []*virtv1.VirtualMachineInstance) []*virtv1.VirtualMachineInstance {
	func (c *VMController) listVMIsFromNamespace(namespace string) ([]*virtv1.VirtualMachineInstance, error) {
	func (c *VMController) listControllerFromNamespace(namespace string) ([]*virtv1.VirtualMachine, error) {
	func (c *VMController) getMatchingControllers(vmi *virtv1.VirtualMachineInstance) (vms []*virtv1.VirtualMachine) {
	func (c *VMController) addVirtualMachine(obj interface{}) {
	func (c *VMController) updateVirtualMachine(old, cur interface{}) {
	func (c *VMController) deleteVirtualMachine(obj interface{}) {
	func (c *VMController) addDataVolume(obj interface{}) {
	func (c *VMController) updateDataVolume(old, cur interface{}) {
	func (c *VMController) deleteDataVolume(obj interface{}) {
	func (c *VMController) addVm(obj interface{}) {
	func (c *VMController) deleteVm(obj interface{}) {
	func (c *VMController) updateVm(old, curr interface{}) {
	func (c *VMController) enqueueVm(obj interface{}) {
	func (c *VMController) updateStatus(vmOrig *virtv1.VirtualMachine, vmi *virtv1.VirtualMachineInstance, createErr error) error {
	func (c *VMController) syncReadyConditionFromVMI(vm *virtv1.VirtualMachine, vmi *virtv1.VirtualMachineInstance) {
	func copyConditionDetails(source *virtv1.VirtualMachineInstanceCondition, dest *virtv1.VirtualMachineCondition) {
	func (c *VMController) processFailure(vm *virtv1.VirtualMachine, vmi *virtv1.VirtualMachineInstance, createErr error) {
	func (c *VMController) resolveControllerRef(namespace string, controllerRef *v1.OwnerReference) *virtv1.VirtualMachine {

replicaset.go   =====================================
	func NewVMIReplicaSet(vmiInformer cache.SharedIndexInformer, vmiRSInformer cache.SharedIndexInformer, recorder record.EventRecorder, clientset kubecli.KubevirtClient, burstReplicas uint) *VMIReplicaSet {
	func (c *VMIReplicaSet) Run(threadiness int, stopCh <-chan struct{}) {
	func (c *VMIReplicaSet) runWorker() {
	func (c *VMIReplicaSet) Execute() bool {
	func (c *VMIReplicaSet) execute(key string) error {
	func (c *VMIReplicaSet) orphan(cm *controller.VirtualMachineControllerRefManager, rs *virtv1.VirtualMachineInstanceReplicaSet, vmis []*virtv1.VirtualMachineInstance) error {
	func (c *VMIReplicaSet) scale(rs *virtv1.VirtualMachineInstanceReplicaSet, vmis []*virtv1.VirtualMachineInstance) error {
	func (c *VMIReplicaSet) filterActiveVMIs(vmis []*virtv1.VirtualMachineInstance) []*virtv1.VirtualMachineInstance {
	func (c *VMIReplicaSet) filterReadyVMIs(vmis []*virtv1.VirtualMachineInstance) []*virtv1.VirtualMachineInstance {
	func (c *VMIReplicaSet) filterFinishedVMIs(vmis []*virtv1.VirtualMachineInstance) []*virtv1.VirtualMachineInstance {
	func filter(vmis []*virtv1.VirtualMachineInstance, f func(vmi *virtv1.VirtualMachineInstance) bool) []*virtv1.VirtualMachineInstance {
	func (c *VMIReplicaSet) listVMIsFromNamespace(namespace string) ([]*virtv1.VirtualMachineInstance, error) {
	func (c *VMIReplicaSet) listControllerFromNamespace(namespace string) ([]*virtv1.VirtualMachineInstanceReplicaSet, error) {
	func (c *VMIReplicaSet) getMatchingControllers(vmi *virtv1.VirtualMachineInstance) (rss []*virtv1.VirtualMachineInstanceReplicaSet) {
	func (c *VMIReplicaSet) addVirtualMachine(obj interface{}) {
	func (c *VMIReplicaSet) updateVirtualMachine(old, cur interface{}) {
	func (c *VMIReplicaSet) deleteVirtualMachine(obj interface{}) {
	func (c *VMIReplicaSet) addReplicaSet(obj interface{}) {
	func (c *VMIReplicaSet) deleteReplicaSet(obj interface{}) {
	func (c *VMIReplicaSet) updateReplicaSet(old, curr interface{}) {
	func (c *VMIReplicaSet) enqueueReplicaSet(obj interface{}) {
	func abs(x int) int {
	func min(x int, y int) int {
	func max(x int, y int) int {
	func limit(x int, burstReplicas uint) int {
	func (c *VMIReplicaSet) hasCondition(rs *virtv1.VirtualMachineInstanceReplicaSet, cond virtv1.VirtualMachineInstanceReplicaSetConditionType) bool {
	func (c *VMIReplicaSet) removeCondition(rs *virtv1.VirtualMachineInstanceReplicaSet, cond virtv1.VirtualMachineInstanceReplicaSetConditionType) {
	func (c *VMIReplicaSet) updateStatus(rs *virtv1.VirtualMachineInstanceReplicaSet, vmis []*virtv1.VirtualMachineInstance, scaleErr error) error {
	func (c *VMIReplicaSet) calcDiff(rs *virtv1.VirtualMachineInstanceReplicaSet, vmis []*virtv1.VirtualMachineInstance) int {
	func (c *VMIReplicaSet) getVirtualMachineBaseName(replicaset *virtv1.VirtualMachineInstanceReplicaSet) string {
	func (c *VMIReplicaSet) checkPaused(rs *virtv1.VirtualMachineInstanceReplicaSet) {
	func (c *VMIReplicaSet) checkFailure(rs *virtv1.VirtualMachineInstanceReplicaSet, diff int, scaleErr error) {
	func OwnerRef(rs *virtv1.VirtualMachineInstanceReplicaSet) metav1.OwnerReference {
	func (c *VMIReplicaSet) resolveControllerRef(namespace string, controllerRef *metav1.OwnerReference) *virtv1.VirtualMachineInstanceReplicaSet {
	func (c *VMIReplicaSet) cleanFinishedVmis(rs *virtv1.VirtualMachineInstanceReplicaSet, vmis []*virtv1.VirtualMachineInstance) error {

node.go   =====================================
	func NewNodeController(clientset kubecli.KubevirtClient, nodeInformer cache.SharedIndexInformer, vmiInformer cache.SharedIndexInformer, recorder record.EventRecorder) *NodeController {
	func (c *NodeController) addNode(obj interface{}) {
	func (c *NodeController) deleteNode(obj interface{}) {
	func (c *NodeController) updateNode(old, curr interface{}) {
	func (c *NodeController) enqueueNode(obj interface{}) {
	func (c *NodeController) addVirtualMachine(obj interface{}) {
	func (c *NodeController) updateVirtualMachine(old, curr interface{}) {
	func (c *NodeController) Run(threadiness int, stopCh <-chan struct{}) {
	func (c *NodeController) runWorker() {
	func (c *NodeController) Execute() bool {
	func (c *NodeController) execute(key string) error {
	func (c *NodeController) alivePodsOnNode(nodeName string) ([]*v1.Pod, error) {
	func filterStuckVirtualMachinesWithoutPods(vmis []*virtv1.VirtualMachineInstance, pods []*v1.Pod) []*virtv1.VirtualMachineInstance {
	func isControlledByVMI(controllerRef *metav1.OwnerReference) bool {
	func isNodeUnresponsive(node *v1.Node, timeout time.Duration) (bool, error) {

migration.go   =====================================
	func NewMigrationController(templateService services.TemplateService,
	func (c *MigrationController) Run(threadiness int, stopCh <-chan struct{}) {
	func (c *MigrationController) runWorker() {
	func (c *MigrationController) Execute() bool {
	func (c *MigrationController) execute(key string) error {
	func (c *MigrationController) canMigrateVMI(migration *virtv1.VirtualMachineInstanceMigration, vmi *virtv1.VirtualMachineInstance) (bool, error) {
	func (c *MigrationController) updateStatus(migration *virtv1.VirtualMachineInstanceMigration, vmi *virtv1.VirtualMachineInstance, pods []*k8sv1.Pod, syncErr error) error {
	func (c *MigrationController) createTargetPod(migration *virtv1.VirtualMachineInstanceMigration, vmi *virtv1.VirtualMachineInstance) error {
	func (c *MigrationController) sync(key string, migration *virtv1.VirtualMachineInstanceMigration, vmi *virtv1.VirtualMachineInstance, pods []*k8sv1.Pod) error {
	func (c *MigrationController) listMatchingTargetPods(migration *virtv1.VirtualMachineInstanceMigration, vmi *virtv1.VirtualMachineInstance) ([]*k8sv1.Pod, error) {
	func (c *MigrationController) addMigration(obj interface{}) {
	func (c *MigrationController) deleteMigration(obj interface{}) {
	func (c *MigrationController) updateMigration(old, curr interface{}) {
	func (c *MigrationController) enqueueMigration(obj interface{}) {
	func (c *MigrationController) getControllerOf(pod *k8sv1.Pod) *v1.OwnerReference {
	func (c *MigrationController) resolveControllerRef(namespace string, controllerRef *v1.OwnerReference) *virtv1.VirtualMachineInstanceMigration {
	func (c *MigrationController) addPod(obj interface{}) {
	func (c *MigrationController) updatePod(old, cur interface{}) {
	func (c *MigrationController) deletePod(obj interface{}) {
	func (c *MigrationController) listMigrationsMatchingVMI(namespace string, name string) ([]*virtv1.VirtualMachineInstanceMigration, error) {
	func (c *MigrationController) addVMI(obj interface{}) {
	func (c *MigrationController) updateVMI(old, cur interface{}) {
	func (c *MigrationController) deleteVMI(obj interface{}) {
	func (c *MigrationController) outboundMigrationsOnNode(node string, runningMigrations []*virtv1.VirtualMachineInstanceMigration) (int, error) {
	func (c *MigrationController) findRunningMigrations() ([]*virtv1.VirtualMachineInstanceMigration, error) {

application.go   =====================================
	func init() {
	func Execute() {
	func (vca *VirtControllerApp) configModificationCallback() {
	func (vca *VirtControllerApp) Run() {
	func (vca *VirtControllerApp) getNewRecorder(namespace string, componentName string) record.EventRecorder {
	func (vca *VirtControllerApp) initCommon() {
	func (vca *VirtControllerApp) initReplicaSet() {
	func (vca *VirtControllerApp) initVirtualMachines() {
	func (vca *VirtControllerApp) initDisruptionBudgetController() {
	func (vca *VirtControllerApp) initEvacuationController() {
	func (vca *VirtControllerApp) initSnapshotController() {
	func (vca *VirtControllerApp) initRestoreController() {
	func (vca *VirtControllerApp) leaderProbe(_ *restful.Request, response *restful.Response) {
	func (vca *VirtControllerApp) AddFlags() {
