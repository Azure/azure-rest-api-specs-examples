
/**
 * Samples for ResourcePools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/
     * GetResourcePool.json
     */
    /**
     * Sample code: GetResourcePool.
     * 
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getResourcePool(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.resourcePools().getByResourceGroupWithResponse("testrg", "HRPool", com.azure.core.util.Context.NONE);
    }
}
