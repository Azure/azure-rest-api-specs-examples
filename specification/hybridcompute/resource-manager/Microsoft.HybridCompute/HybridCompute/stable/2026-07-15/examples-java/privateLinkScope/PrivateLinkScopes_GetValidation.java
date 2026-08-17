
/**
 * Samples for PrivateLinkScopes GetValidationDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/privateLinkScope/PrivateLinkScopes_GetValidation.json
     */
    /**
     * Sample code: PrivateLinkScopeGet.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void privateLinkScopeGet(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateLinkScopes().getValidationDetailsWithResponse("wus2", "f5dc51d3-92ed-4d7e-947a-775ea79b4919",
            com.azure.core.util.Context.NONE);
    }
}
