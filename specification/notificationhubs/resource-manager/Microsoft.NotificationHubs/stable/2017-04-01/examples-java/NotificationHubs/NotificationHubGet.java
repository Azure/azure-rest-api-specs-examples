import com.azure.core.util.Context;

/** Samples for NotificationHubs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubGet.json
     */
    /**
     * Sample code: NotificationHubGet.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubGet(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().getWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub", Context.NONE);
    }
}
