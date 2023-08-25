/** Samples for Notification ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListNotifications.json
     */
    /**
     * Sample code: ApiManagementListNotifications.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListNotifications(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.notifications().listByService("rg1", "apimService1", null, null, com.azure.core.util.Context.NONE);
    }
}
