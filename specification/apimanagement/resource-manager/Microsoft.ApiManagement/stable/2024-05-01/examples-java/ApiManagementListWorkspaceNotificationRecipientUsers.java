
import com.azure.resourcemanager.apimanagement.models.NotificationName;

/**
 * Samples for WorkspaceNotificationRecipientUser ListByNotification.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceNotificationRecipientUsers.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceNotificationRecipientUsers.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListWorkspaceNotificationRecipientUsers(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNotificationRecipientUsers().listByNotificationWithResponse("rg1", "apimService1", "wks1",
            NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE, com.azure.core.util.Context.NONE);
    }
}
