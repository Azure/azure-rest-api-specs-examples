import com.azure.resourcemanager.apimanagement.models.NotificationName;

/** Samples for NotificationRecipientEmail ListByNotification. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListNotificationRecipientEmails.json
     */
    /**
     * Sample code: ApiManagementListNotificationRecipientEmails.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListNotificationRecipientEmails(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .notificationRecipientEmails()
            .listByNotificationWithResponse(
                "rg1",
                "apimService1",
                NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE,
                com.azure.core.util.Context.NONE);
    }
}
