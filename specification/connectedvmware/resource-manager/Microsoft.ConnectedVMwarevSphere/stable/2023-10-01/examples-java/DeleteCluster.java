/** Samples for Clusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteCluster.json
     */
    /**
     * Sample code: DeleteCluster.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteCluster(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().delete("testrg", "HRCluster", null, com.azure.core.util.Context.NONE);
    }
}
