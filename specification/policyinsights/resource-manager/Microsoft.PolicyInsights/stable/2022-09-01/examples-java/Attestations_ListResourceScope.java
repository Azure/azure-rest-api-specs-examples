
/**
 * Samples for Attestations ListForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/
     * Attestations_ListResourceScope.json
     */
    /**
     * Sample code: List attestations at individual resource scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listAttestationsAtIndividualResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.attestations().listForResource(
            "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myrg/providers/microsoft.compute/virtualMachines/devVM",
            null, null, com.azure.core.util.Context.NONE);
    }
}
