
/**
 * Samples for BillingStatistics Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/billingStatistics/GetBillingStatistic.json
     */
    /**
     * Sample code: Get a billing statistic.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getABillingStatistic(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.billingStatistics().getWithResponse("myRg", "myWorkspace", "sapSolutionUsage",
            com.azure.core.util.Context.NONE);
    }
}
