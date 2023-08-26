import com.azure.resourcemanager.apimanagement.models.NotificationName;

/** Samples for NotificationRecipientEmail CheckEntityExists. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadNotificationRecipientEmail.json
     */
    /**
     * Sample code: ApiManagementHeadNotificationRecipientEmail.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadNotificationRecipientEmail(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .notificationRecipientEmails()
            .checkEntityExistsWithResponse(
                "rg1",
                "apimService1",
                NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE,
                "contoso@live.com",
                com.azure.core.util.Context.NONE);
    }
}
