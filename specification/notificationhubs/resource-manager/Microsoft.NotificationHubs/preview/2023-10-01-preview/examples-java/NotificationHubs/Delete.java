
/**
 * Samples for NotificationHubs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/Delete.json
     */
    /**
     * Sample code: NotificationHubs_Delete.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        notificationHubsDelete(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().deleteWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub",
            com.azure.core.util.Context.NONE);
    }
}
