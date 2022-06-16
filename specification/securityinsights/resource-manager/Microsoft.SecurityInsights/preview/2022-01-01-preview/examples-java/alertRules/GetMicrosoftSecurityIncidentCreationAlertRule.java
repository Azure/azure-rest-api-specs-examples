import com.azure.core.util.Context;

/** Samples for AlertRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/alertRules/GetMicrosoftSecurityIncidentCreationAlertRule.json
     */
    /**
     * Sample code: Get a MicrosoftSecurityIncidentCreation rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAMicrosoftSecurityIncidentCreationRule(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .alertRules()
            .getWithResponse("myRg", "myWorkspace", "microsoftSecurityIncidentCreationRuleExample", Context.NONE);
    }
}
