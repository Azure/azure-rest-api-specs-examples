import com.azure.core.util.Context;

/** Samples for Subscriptions ListQuotas. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Subscription_ListQuotas.json
     */
    /**
     * Sample code: List subscription quota information in West US.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listSubscriptionQuotaInformationInWestUS(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.subscriptions().listQuotasWithResponse("West US", Context.NONE);
    }
}
