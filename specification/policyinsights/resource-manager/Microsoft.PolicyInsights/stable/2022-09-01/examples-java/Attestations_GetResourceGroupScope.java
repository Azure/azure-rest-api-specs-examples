import com.azure.core.util.Context;

/** Samples for Attestations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/Attestations_GetResourceGroupScope.json
     */
    /**
     * Sample code: Get attestation at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void getAttestationAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .attestations()
            .getByResourceGroupWithResponse("myRg", "790996e6-9871-4b1f-9cd9-ec42cd6ced1e", Context.NONE);
    }
}
