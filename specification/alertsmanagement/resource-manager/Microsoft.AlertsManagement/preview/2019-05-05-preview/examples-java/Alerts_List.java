import com.azure.core.util.Context;

/** Samples for Alerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_List.json
     */
    /**
     * Sample code: ListAlerts.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void listAlerts(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager
            .alerts()
            .list(
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
