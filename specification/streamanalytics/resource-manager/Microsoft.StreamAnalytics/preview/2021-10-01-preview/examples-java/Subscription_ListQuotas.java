
/**
 * Samples for Subscriptions ListQuotas.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Subscription_ListQuotas.json
     */
    /**
     * Sample code: List subscription quota information in West US.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listSubscriptionQuotaInformationInWestUS(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.subscriptions().listQuotasWithResponse("West US", com.azure.core.util.Context.NONE);
    }
}
