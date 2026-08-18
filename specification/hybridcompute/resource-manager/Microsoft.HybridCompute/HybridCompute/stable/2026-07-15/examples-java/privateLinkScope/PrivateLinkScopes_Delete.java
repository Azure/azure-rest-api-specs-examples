
/**
 * Samples for PrivateLinkScopes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/privateLinkScope/PrivateLinkScopes_Delete.json
     */
    /**
     * Sample code: PrivateLinkScopesDelete.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void privateLinkScopesDelete(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateLinkScopes().delete("my-resource-group", "my-privatelinkscope",
            com.azure.core.util.Context.NONE);
    }
}
