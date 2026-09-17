
/**
 * Samples for AlertRuleTemplates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/alertRuleTemplates/GetAlertRuleTemplates.json
     */
    /**
     * Sample code: Get all alert rule templates.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllAlertRuleTemplates(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRuleTemplates().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
