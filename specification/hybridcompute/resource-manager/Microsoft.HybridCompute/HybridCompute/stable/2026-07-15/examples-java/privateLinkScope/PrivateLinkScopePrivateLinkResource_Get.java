
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/privateLinkScope/PrivateLinkScopePrivateLinkResource_Get.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        getsPrivateEndpointConnection(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateLinkResources().getWithResponse("myResourceGroup", "myPrivateLinkScope", "hybridcompute",
            com.azure.core.util.Context.NONE);
    }
}
