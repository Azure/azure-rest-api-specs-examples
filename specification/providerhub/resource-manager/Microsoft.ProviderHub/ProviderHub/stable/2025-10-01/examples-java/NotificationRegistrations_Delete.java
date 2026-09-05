
/**
 * Samples for NotificationRegistrations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/NotificationRegistrations_Delete.json
     */
    /**
     * Sample code: NotificationRegistrations_Delete.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        notificationRegistrationsDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.notificationRegistrations().deleteByResourceGroupWithResponse("Microsoft.Contoso",
            "fooNotificationRegistration", com.azure.core.util.Context.NONE);
    }
}
