
/**
 * Samples for PrivateLinkScopes GetValidationDetailsForMachine.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/privateLinkScope/PrivateLinkScopes_GetValidationForMachine.json
     */
    /**
     * Sample code: PrivateLinkScopeGet.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void privateLinkScopeGet(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateLinkScopes().getValidationDetailsForMachineWithResponse("my-resource-group", "machineName",
            com.azure.core.util.Context.NONE);
    }
}
