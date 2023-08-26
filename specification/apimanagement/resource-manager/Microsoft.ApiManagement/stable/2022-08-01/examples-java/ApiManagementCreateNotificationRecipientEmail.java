import com.azure.resourcemanager.apimanagement.models.NotificationName;

/** Samples for NotificationRecipientEmail CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateNotificationRecipientEmail.json
     */
    /**
     * Sample code: ApiManagementCreateNotificationRecipientEmail.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateNotificationRecipientEmail(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .notificationRecipientEmails()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                NotificationName.REQUEST_PUBLISHER_NOTIFICATION_MESSAGE,
                "foobar@live.com",
                com.azure.core.util.Context.NONE);
    }
}
