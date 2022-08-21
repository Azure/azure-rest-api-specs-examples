import com.azure.core.util.Context;

/** Samples for Datastores GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetDatastore.json
     */
    /**
     * Sample code: GetDatastore.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getDatastore(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.datastores().getByResourceGroupWithResponse("testrg", "HRDatastore", Context.NONE);
    }
}
