
import com.azure.resourcemanager.apimanagement.models.NotificationName;

/**
 * Samples for WorkspaceNotification CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceNotification.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceNotification.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceNotification(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNotifications().createOrUpdateWithResponse("rg1", "apimService1", "wks1",
            NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE, null, com.azure.core.util.Context.NONE);
    }
}
