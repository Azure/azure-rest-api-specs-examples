import com.azure.core.util.Context;

/** Samples for Attestations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_ListSubscriptionScope.json
     */
    /**
     * Sample code: List attestations at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listAttestationsAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.attestations().list(null, null, Context.NONE);
    }
}
