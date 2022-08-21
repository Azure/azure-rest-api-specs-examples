import com.azure.core.util.Context;

/** Samples for ResourcePools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteResourcePool.json
     */
    /**
     * Sample code: DeleteResourcePool.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteResourcePool(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.resourcePools().delete("testrg", "HRPool", null, Context.NONE);
    }
}
