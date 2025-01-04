
/**
 * Samples for SecurityMLAnalyticsSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/
     * securityMLAnalyticsSettings/DeleteSecurityMLAnalyticsSetting.json
     */
    /**
     * Sample code: Delete a Security ML Analytics Settings.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteASecurityMLAnalyticsSettings(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.securityMLAnalyticsSettings().deleteWithResponse("myRg", "myWorkspace",
            "f209187f-1d17-4431-94af-c141bf5f23db", com.azure.core.util.Context.NONE);
    }
}
