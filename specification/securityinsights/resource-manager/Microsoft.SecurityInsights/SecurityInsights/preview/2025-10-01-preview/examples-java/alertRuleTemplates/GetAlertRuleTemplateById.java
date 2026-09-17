
/**
 * Samples for AlertRuleTemplates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/alertRuleTemplates/GetAlertRuleTemplateById.json
     */
    /**
     * Sample code: Get alert rule template by Id.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAlertRuleTemplateById(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRuleTemplates().getWithResponse("myRg", "myWorkspace", "65360bb0-8986-4ade-a89d-af3cf44d28aa",
            com.azure.core.util.Context.NONE);
    }
}
