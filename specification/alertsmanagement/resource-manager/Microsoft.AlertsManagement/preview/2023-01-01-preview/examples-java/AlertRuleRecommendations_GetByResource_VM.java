
/**
 * Samples for AlertRuleRecommendations ListByResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-01-01-preview/examples/
     * AlertRuleRecommendations_GetByResource_VM.json
     */
    /**
     * Sample code: List alert rule recommendations for virtual machines at resource level.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void listAlertRuleRecommendationsForVirtualMachinesAtResourceLevel(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertRuleRecommendations().listByResource(
            "subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourcegroups/test/providers/Microsoft.Compute/virtualMachines/testMachineCanBeSafelyDeleted",
            com.azure.core.util.Context.NONE);
    }
}
