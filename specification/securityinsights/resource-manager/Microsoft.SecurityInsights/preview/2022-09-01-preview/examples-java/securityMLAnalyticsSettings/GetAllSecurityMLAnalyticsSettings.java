import com.azure.core.util.Context;

/** Samples for SecurityMLAnalyticsSettings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/securityMLAnalyticsSettings/GetAllSecurityMLAnalyticsSettings.json
     */
    /**
     * Sample code: Get all Security ML Analytics Settings.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllSecurityMLAnalyticsSettings(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.securityMLAnalyticsSettings().list("myRg", "myWorkspace", Context.NONE);
    }
}
