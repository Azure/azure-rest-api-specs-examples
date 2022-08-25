import com.azure.core.util.Context;

/** Samples for Alerts GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_GetById.json
     */
    /**
     * Sample code: GetById.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void getById(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alerts().getByIdWithResponse("66114d64-d9d9-478b-95c9-b789d6502100", Context.NONE);
    }
}
