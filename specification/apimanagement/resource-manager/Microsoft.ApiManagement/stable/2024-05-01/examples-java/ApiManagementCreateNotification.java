
import com.azure.resourcemanager.apimanagement.models.NotificationName;

/**
 * Samples for Notification CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateNotification.json
     */
    /**
     * Sample code: ApiManagementCreateNotification.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateNotification(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.notifications().createOrUpdateWithResponse("rg1", "apimService1",
            NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE, null, com.azure.core.util.Context.NONE);
    }
}
