import com.azure.core.util.Context;

/** Samples for Actions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/actions/DeleteActionOfAlertRule.json
     */
    /**
     * Sample code: Delete an action of alert rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnActionOfAlertRule(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .actions()
            .deleteWithResponse(
                "myRg",
                "myWorkspace",
                "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                "912bec42-cb66-4c03-ac63-1761b6898c3e",
                Context.NONE);
    }
}
