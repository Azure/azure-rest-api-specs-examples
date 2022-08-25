import com.azure.core.util.Context;

/** Samples for AlertProcessingRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_List_Subscription.json
     */
    /**
     * Sample code: GetAlertProcessingRulesSubscriptionWide.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void getAlertProcessingRulesSubscriptionWide(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().list(Context.NONE);
    }
}
