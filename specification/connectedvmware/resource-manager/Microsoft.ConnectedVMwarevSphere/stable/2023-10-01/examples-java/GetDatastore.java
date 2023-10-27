/** Samples for Datastores GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/GetDatastore.json
     */
    /**
     * Sample code: GetDatastore.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getDatastore(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.datastores().getByResourceGroupWithResponse("testrg", "HRDatastore", com.azure.core.util.Context.NONE);
    }
}
