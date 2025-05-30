
import com.azure.resourcemanager.apimanagement.models.NotificationName;

/**
 * Samples for WorkspaceNotificationRecipientUser CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceNotificationRecipientUser.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceNotificationRecipientUser.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspaceNotificationRecipientUser(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNotificationRecipientUsers().createOrUpdateWithResponse("rg1", "apimService1", "wks1",
            NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE, "576823d0a40f7e74ec07d642",
            com.azure.core.util.Context.NONE);
    }
}
