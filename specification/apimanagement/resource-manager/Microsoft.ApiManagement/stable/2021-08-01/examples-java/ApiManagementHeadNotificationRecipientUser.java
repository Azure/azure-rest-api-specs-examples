import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.NotificationName;

/** Samples for NotificationRecipientUser CheckEntityExists. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadNotificationRecipientUser.json
     */
    /**
     * Sample code: ApiManagementHeadNotificationRecipientUser.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadNotificationRecipientUser(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .notificationRecipientUsers()
            .checkEntityExistsWithResponse(
                "rg1",
                "apimService1",
                NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE,
                "576823d0a40f7e74ec07d642",
                Context.NONE);
    }
}
