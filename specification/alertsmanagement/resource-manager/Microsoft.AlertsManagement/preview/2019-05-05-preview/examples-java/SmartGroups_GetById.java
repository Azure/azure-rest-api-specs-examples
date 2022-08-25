import com.azure.core.util.Context;

/** Samples for SmartGroups GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_GetById.json
     */
    /**
     * Sample code: Get.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void get(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.smartGroups().getByIdWithResponse("603675da-9851-4b26-854a-49fc53d32715", Context.NONE);
    }
}
