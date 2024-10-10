
/**
 * Samples for AlertRuleRecommendations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-01-01-preview/examples/
     * AlertRuleRecommendations_GetBySubscription_VM.json
     */
    /**
     * Sample code: List alert rule recommendations for virtual machines at subscription level.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void listAlertRuleRecommendationsForVirtualMachinesAtSubscriptionLevel(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertRuleRecommendations().list("microsoft.compute/virtualmachines", com.azure.core.util.Context.NONE);
    }
}
