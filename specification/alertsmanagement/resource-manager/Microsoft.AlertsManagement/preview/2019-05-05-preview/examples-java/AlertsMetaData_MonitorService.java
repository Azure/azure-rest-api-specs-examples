import com.azure.core.util.Context;
import com.azure.resourcemanager.alertsmanagement.models.Identifier;

/** Samples for Alerts Metadata. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/AlertsMetaData_MonitorService.json
     */
    /**
     * Sample code: MonService.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void monService(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alerts().metadataWithResponse(Identifier.MONITOR_SERVICE_LIST, Context.NONE);
    }
}
