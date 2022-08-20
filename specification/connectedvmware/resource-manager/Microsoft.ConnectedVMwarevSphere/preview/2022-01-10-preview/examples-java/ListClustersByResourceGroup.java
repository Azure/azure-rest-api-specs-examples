import com.azure.core.util.Context;

/** Samples for Clusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListClustersByResourceGroup.json
     */
    /**
     * Sample code: ListClustersByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listClustersByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().listByResourceGroup("testrg", Context.NONE);
    }
}
