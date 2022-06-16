/** Samples for NotificationHubs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubCreate.json
     */
    /**
     * Sample code: NotificationHubCreate.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubCreate(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .notificationHubs()
            .define("nh-sdk-hub")
            .withRegion("eastus")
            .withExistingNamespace("5ktrial", "nh-sdk-ns")
            .create();
    }
}
