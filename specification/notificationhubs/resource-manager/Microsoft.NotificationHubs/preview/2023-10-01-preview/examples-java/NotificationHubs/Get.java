
/**
 * Samples for NotificationHubs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/Get.json
     */
    /**
     * Sample code: NotificationHubs_Get.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubsGet(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().getWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub",
            com.azure.core.util.Context.NONE);
    }
}
