
/**
 * Samples for PrivateLinkResources ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/privateLinkScope/PrivateLinkScopePrivateLinkResource_ListGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        getsPrivateEndpointConnection(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateLinkResources().listByPrivateLinkScope("myResourceGroup", "myPrivateLinkScope",
            com.azure.core.util.Context.NONE);
    }
}
