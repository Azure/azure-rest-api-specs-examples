import com.azure.core.util.Context;

/** Samples for ResourcePools GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetResourcePool.json
     */
    /**
     * Sample code: GetResourcePool.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getResourcePool(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.resourcePools().getByResourceGroupWithResponse("testrg", "HRPool", Context.NONE);
    }
}
