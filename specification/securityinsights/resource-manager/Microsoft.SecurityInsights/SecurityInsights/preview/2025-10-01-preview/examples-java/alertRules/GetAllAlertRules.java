
/**
 * Samples for AlertRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/alertRules/GetAllAlertRules.json
     */
    /**
     * Sample code: Get all alert rules.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllAlertRules(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRules().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
