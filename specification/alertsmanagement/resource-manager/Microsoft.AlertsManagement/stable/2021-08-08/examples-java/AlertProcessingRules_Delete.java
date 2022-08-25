import com.azure.core.util.Context;

/** Samples for AlertProcessingRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Delete.json
     */
    /**
     * Sample code: DeleteAlertProcessingRule.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void deleteAlertProcessingRule(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().deleteWithResponse("alertscorrelationrg", "DailySuppression", Context.NONE);
    }
}
