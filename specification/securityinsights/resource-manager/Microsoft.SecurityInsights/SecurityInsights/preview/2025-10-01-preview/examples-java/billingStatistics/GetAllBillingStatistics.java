
/**
 * Samples for BillingStatistics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/billingStatistics/GetAllBillingStatistics.json
     */
    /**
     * Sample code: Get all Microsoft Sentinel billing statistics.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllMicrosoftSentinelBillingStatistics(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.billingStatistics().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
