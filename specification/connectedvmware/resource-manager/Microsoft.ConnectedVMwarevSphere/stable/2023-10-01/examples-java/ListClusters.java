/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListClusters.json
     */
    /**
     * Sample code: ListClusters.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listClusters(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}
