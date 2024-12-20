
import com.azure.resourcemanager.alertsmanagement.models.AlertsSummaryGroupByFields;

/**
 * Samples for Alerts GetSummary.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/
     * Alerts_Summary.json
     */
    /**
     * Sample code: Summary.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void summary(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alerts().getSummaryWithResponse(AlertsSummaryGroupByFields.fromString("severity,alertState"), null,
            null, null, null, null, null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
