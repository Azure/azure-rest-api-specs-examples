import com.azure.resourcemanager.apimanagement.models.NotificationName;

/** Samples for NotificationRecipientUser ListByNotification. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListNotificationRecipientUsers.json
     */
    /**
     * Sample code: ApiManagementListNotificationRecipientUsers.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListNotificationRecipientUsers(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .notificationRecipientUsers()
            .listByNotificationWithResponse(
                "rg1",
                "apimService1",
                NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE,
                com.azure.core.util.Context.NONE);
    }
}
