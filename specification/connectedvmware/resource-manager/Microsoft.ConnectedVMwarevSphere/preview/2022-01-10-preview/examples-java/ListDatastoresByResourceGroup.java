import com.azure.core.util.Context;

/** Samples for Datastores ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListDatastoresByResourceGroup.json
     */
    /**
     * Sample code: ListDatastoresByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listDatastoresByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.datastores().listByResourceGroup("testrg", Context.NONE);
    }
}
