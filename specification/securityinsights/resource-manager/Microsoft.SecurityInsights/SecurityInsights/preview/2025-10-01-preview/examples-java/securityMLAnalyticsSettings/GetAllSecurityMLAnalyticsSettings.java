
/**
 * Samples for SecurityMLAnalyticsSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/securityMLAnalyticsSettings/GetAllSecurityMLAnalyticsSettings.json
     */
    /**
     * Sample code: Get all Security ML Analytics Settings.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllSecurityMLAnalyticsSettings(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.securityMLAnalyticsSettings().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
