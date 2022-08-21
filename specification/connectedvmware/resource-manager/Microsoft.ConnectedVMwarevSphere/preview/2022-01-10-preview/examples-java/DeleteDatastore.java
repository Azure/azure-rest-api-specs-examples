import com.azure.core.util.Context;

/** Samples for Datastores Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteDatastore.json
     */
    /**
     * Sample code: DeleteDatastore.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteDatastore(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.datastores().delete("testrg", "HRDatastore", null, Context.NONE);
    }
}
