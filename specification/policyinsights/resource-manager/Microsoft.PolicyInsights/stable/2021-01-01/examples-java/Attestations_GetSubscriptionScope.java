import com.azure.core.util.Context;

/** Samples for Attestations GetAtSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_GetSubscriptionScope.json
     */
    /**
     * Sample code: Get attestation at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void getAttestationAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.attestations().getAtSubscriptionWithResponse("790996e6-9871-4b1f-9cd9-ec42cd6ced1e", Context.NONE);
    }
}
