/** Samples for NotificationRegistrations ListByProviderRegistration. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/NotificationRegistrations_ListByProviderRegistration.json
     */
    /**
     * Sample code: NotificationRegistrations_ListByProviderRegistration.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void notificationRegistrationsListByProviderRegistration(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .notificationRegistrations()
            .listByProviderRegistration("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
