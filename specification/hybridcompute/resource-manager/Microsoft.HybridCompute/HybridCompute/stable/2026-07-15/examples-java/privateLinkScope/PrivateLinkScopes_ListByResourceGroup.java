
/**
 * Samples for PrivateLinkScopes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/privateLinkScope/PrivateLinkScopes_ListByResourceGroup.json
     */
    /**
     * Sample code: PrivateLinkScopeListByResourceGroup.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        privateLinkScopeListByResourceGroup(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateLinkScopes().listByResourceGroup("my-resource-group", com.azure.core.util.Context.NONE);
    }
}
