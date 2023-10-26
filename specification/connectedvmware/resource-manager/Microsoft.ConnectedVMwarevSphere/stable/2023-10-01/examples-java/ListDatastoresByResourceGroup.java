/** Samples for Datastores ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListDatastoresByResourceGroup.json
     */
    /**
     * Sample code: ListDatastoresByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listDatastoresByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.datastores().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
