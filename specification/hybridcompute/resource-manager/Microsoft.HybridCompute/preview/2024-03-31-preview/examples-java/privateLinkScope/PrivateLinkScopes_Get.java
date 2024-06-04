
/**
 * Samples for PrivateLinkScopes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/
     * privateLinkScope/PrivateLinkScopes_Get.json
     */
    /**
     * Sample code: PrivateLinkScopeGet.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void privateLinkScopeGet(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateLinkScopes().getByResourceGroupWithResponse("my-resource-group", "my-privatelinkscope",
            com.azure.core.util.Context.NONE);
    }
}
