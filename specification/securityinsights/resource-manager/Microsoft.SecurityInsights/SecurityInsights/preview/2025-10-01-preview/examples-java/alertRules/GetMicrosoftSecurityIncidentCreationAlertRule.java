
/**
 * Samples for AlertRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/alertRules/GetMicrosoftSecurityIncidentCreationAlertRule.json
     */
    /**
     * Sample code: Get a MicrosoftSecurityIncidentCreation rule.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAMicrosoftSecurityIncidentCreationRule(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRules().getWithResponse("myRg", "myWorkspace", "microsoftSecurityIncidentCreationRuleExample",
            com.azure.core.util.Context.NONE);
    }
}
