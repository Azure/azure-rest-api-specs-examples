
/**
 * Samples for Attestations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/
     * Attestations_ListSubscriptionScope.json
     */
    /**
     * Sample code: List attestations at subscription scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        listAttestationsAtSubscriptionScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.attestations().list(null, null, com.azure.core.util.Context.NONE);
    }
}
