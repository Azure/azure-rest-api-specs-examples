
import com.azure.resourcemanager.apimanagement.models.NotificationName;

/**
 * Samples for WorkspaceNotificationRecipientEmail CheckEntityExists.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceNotificationRecipientEmail.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceNotificationRecipientEmail.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadWorkspaceNotificationRecipientEmail(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNotificationRecipientEmails().checkEntityExistsWithResponse("rg1", "apimService1", "wks1",
            NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE, "contoso@live.com",
            com.azure.core.util.Context.NONE);
    }
}
