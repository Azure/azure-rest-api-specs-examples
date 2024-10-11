
import com.azure.resourcemanager.alertsmanagement.models.AlertState;

/**
 * Samples for SmartGroups ChangeState.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/
     * SmartGroups_ChangeState.json
     */
    /**
     * Sample code: changestate.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void changestate(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.smartGroups().changeStateWithResponse("a808445e-bb38-4751-85c2-1b109ccc1059", AlertState.ACKNOWLEDGED,
            com.azure.core.util.Context.NONE);
    }
}
