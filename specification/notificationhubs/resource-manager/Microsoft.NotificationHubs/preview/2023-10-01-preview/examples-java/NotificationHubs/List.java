
/**
 * Samples for NotificationHubs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/List.json
     */
    /**
     * Sample code: NotificationHubs_List.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        notificationHubsList(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().list("5ktrial", "nh-sdk-ns", null, null, com.azure.core.util.Context.NONE);
    }
}
