
/**
 * Samples for Alerts GetHistory.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/
     * Alerts_History.json
     */
    /**
     * Sample code: Resolve.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void resolve(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alerts().getHistoryWithResponse("66114d64-d9d9-478b-95c9-b789d6502100",
            com.azure.core.util.Context.NONE);
    }
}
