
/**
 * Samples for NotificationRegistrations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/NotificationRegistrations_Get.json
     */
    /**
     * Sample code: NotificationRegistrations_Get.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void notificationRegistrationsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.notificationRegistrations().getWithResponse("Microsoft.Contoso", "fooNotificationRegistration",
            com.azure.core.util.Context.NONE);
    }
}
