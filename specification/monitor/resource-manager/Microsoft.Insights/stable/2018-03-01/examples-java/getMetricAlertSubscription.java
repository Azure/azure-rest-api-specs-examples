
import com.azure.core.util.Context;

/** Samples for MetricAlerts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getMetricAlertSubscription.
     * json
     */
    /**
     * Sample code: Get an alert rule on subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnAlertRuleOnSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest",
            "MetricAlertAtSubscriptionLevel", Context.NONE);
    }
}
