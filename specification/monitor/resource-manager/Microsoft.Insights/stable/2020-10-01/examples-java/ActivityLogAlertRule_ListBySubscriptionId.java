import com.azure.core.util.Context;

/** Samples for ActivityLogAlerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_ListBySubscriptionId.json
     */
    /**
     * Sample code: Get list of all Activity Log Alert rules under a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getListOfAllActivityLogAlertRulesUnderASubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActivityLogAlerts().list(Context.NONE);
    }
}
