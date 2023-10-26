/** Samples for Datastores List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListDatastores.json
     */
    /**
     * Sample code: ListDatastores.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listDatastores(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.datastores().list(com.azure.core.util.Context.NONE);
    }
}
