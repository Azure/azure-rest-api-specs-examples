/** Samples for Clusters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/GetCluster.json
     */
    /**
     * Sample code: GetCluster.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getCluster(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().getByResourceGroupWithResponse("testrg", "HRCluster", com.azure.core.util.Context.NONE);
    }
}
