import com.azure.core.util.Context;

/** Samples for AlertRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/alertRules/GetFusionAlertRule.json
     */
    /**
     * Sample code: Get a Fusion alert rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAFusionAlertRule(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRules().getWithResponse("myRg", "myWorkspace", "myFirstFusionRule", Context.NONE);
    }
}
