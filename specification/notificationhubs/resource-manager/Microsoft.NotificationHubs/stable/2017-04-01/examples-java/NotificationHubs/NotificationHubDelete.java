import com.azure.core.util.Context;

/** Samples for NotificationHubs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubDelete.json
     */
    /**
     * Sample code: NotificationHubDelete.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubDelete(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().deleteWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub", Context.NONE);
    }
}
