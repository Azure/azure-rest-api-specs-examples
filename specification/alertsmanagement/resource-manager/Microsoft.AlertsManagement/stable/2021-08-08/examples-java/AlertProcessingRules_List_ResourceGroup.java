import com.azure.core.util.Context;

/** Samples for AlertProcessingRules ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_List_ResourceGroup.json
     */
    /**
     * Sample code: GetAlertProcessingRulesResourceGroupWide.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void getAlertProcessingRulesResourceGroupWide(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().listByResourceGroup("alertscorrelationrg", Context.NONE);
    }
}
