/** Samples for Clusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListClustersByResourceGroup.json
     */
    /**
     * Sample code: ListClustersByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listClustersByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
