import com.azure.core.util.Context;

/** Samples for AlertRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/alertRules/GetNrtAlertRule.json
     */
    /**
     * Sample code: Get an Nrt alert rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnNrtAlertRule(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .alertRules()
            .getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", Context.NONE);
    }
}
