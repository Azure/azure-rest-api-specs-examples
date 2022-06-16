import com.azure.core.util.Context;

/** Samples for Remediations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListSubscriptionScope.json
     */
    /**
     * Sample code: List remediations at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listRemediationsAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().list(null, null, Context.NONE);
    }
}
