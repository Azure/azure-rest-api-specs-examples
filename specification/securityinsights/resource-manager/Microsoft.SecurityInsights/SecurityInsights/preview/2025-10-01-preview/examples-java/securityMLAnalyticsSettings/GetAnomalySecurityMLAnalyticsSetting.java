
/**
 * Samples for SecurityMLAnalyticsSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/securityMLAnalyticsSettings/GetAnomalySecurityMLAnalyticsSetting.json
     */
    /**
     * Sample code: Get a Anomaly Security ML Analytics Settings.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAAnomalySecurityMLAnalyticsSettings(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.securityMLAnalyticsSettings().getWithResponse("myRg", "myWorkspace", "myFirstAnomalySettings",
            com.azure.core.util.Context.NONE);
    }
}
