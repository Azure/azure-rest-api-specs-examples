/** Samples for Datastores Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteDatastore.json
     */
    /**
     * Sample code: DeleteDatastore.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteDatastore(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.datastores().delete("testrg", "HRDatastore", null, com.azure.core.util.Context.NONE);
    }
}
