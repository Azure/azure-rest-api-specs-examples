
/**
 * Samples for PrivateLinkScopes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/privateLinkScope/PrivateLinkScopes_Get.json
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
