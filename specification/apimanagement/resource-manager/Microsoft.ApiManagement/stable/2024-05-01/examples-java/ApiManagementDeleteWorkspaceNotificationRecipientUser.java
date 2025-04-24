
import com.azure.resourcemanager.apimanagement.models.NotificationName;

/**
 * Samples for WorkspaceNotificationRecipientUser Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceNotificationRecipientUser.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceNotificationRecipientUser.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceNotificationRecipientUser(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNotificationRecipientUsers().deleteWithResponse("rg1", "apimService1", "wks1",
            NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE, "576823d0a40f7e74ec07d642",
            com.azure.core.util.Context.NONE);
    }
}
