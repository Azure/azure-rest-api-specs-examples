import com.azure.core.util.Context;

/** Samples for AlertProcessingRules GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_GetById.json
     */
    /**
     * Sample code: GetAlertProcessingRuleById.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void getAlertProcessingRuleById(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager
            .alertProcessingRules()
            .getByResourceGroupWithResponse("alertscorrelationrg", "DailySuppression", Context.NONE);
    }
}
