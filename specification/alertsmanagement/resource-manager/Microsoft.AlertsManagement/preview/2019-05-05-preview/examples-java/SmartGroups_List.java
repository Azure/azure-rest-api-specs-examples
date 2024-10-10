
/**
 * Samples for SmartGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/
     * SmartGroups_List.json
     */
    /**
     * Sample code: List.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void list(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.smartGroups().list(null, null, null, null, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
