/** Samples for ResourcePools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteResourcePool.json
     */
    /**
     * Sample code: DeleteResourcePool.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteResourcePool(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.resourcePools().delete("testrg", "HRPool", null, com.azure.core.util.Context.NONE);
    }
}
