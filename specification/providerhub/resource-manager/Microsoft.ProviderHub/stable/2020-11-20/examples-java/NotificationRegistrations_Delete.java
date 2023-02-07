/** Samples for NotificationRegistrations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/NotificationRegistrations_Delete.json
     */
    /**
     * Sample code: NotificationRegistrations_Delete.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void notificationRegistrationsDelete(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .notificationRegistrations()
            .deleteByResourceGroupWithResponse(
                "Microsoft.Contoso", "fooNotificationRegistration", com.azure.core.util.Context.NONE);
    }
}
