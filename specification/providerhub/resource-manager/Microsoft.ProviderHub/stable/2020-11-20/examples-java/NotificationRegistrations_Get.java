/** Samples for NotificationRegistrations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/NotificationRegistrations_Get.json
     */
    /**
     * Sample code: NotificationRegistrations_Get.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void notificationRegistrationsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .notificationRegistrations()
            .getWithResponse("Microsoft.Contoso", "fooNotificationRegistration", com.azure.core.util.Context.NONE);
    }
}
