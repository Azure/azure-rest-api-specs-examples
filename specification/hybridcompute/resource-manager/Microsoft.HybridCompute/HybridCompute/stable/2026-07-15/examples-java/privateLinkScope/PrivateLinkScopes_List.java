
/**
 * Samples for PrivateLinkScopes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/privateLinkScope/PrivateLinkScopes_List.json
     */
    /**
     * Sample code: PrivateLinkScopesList.json.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void privateLinkScopesListJson(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateLinkScopes().list(com.azure.core.util.Context.NONE);
    }
}
