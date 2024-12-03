
/**
 * Samples for NotificationHubs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/CreateOrUpdate.json
     */
    /**
     * Sample code: NotificationHubs_CreateOrUpdate.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        notificationHubsCreateOrUpdate(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().define("nh-sdk-hub").withRegion("eastus")
            .withExistingNamespace("5ktrial", "nh-sdk-ns").create();
    }
}
