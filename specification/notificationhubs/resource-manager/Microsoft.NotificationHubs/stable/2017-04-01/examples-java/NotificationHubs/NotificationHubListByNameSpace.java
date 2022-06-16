import com.azure.core.util.Context;

/** Samples for NotificationHubs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubListByNameSpace.json
     */
    /**
     * Sample code: NotificationHubListByNameSpace.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubListByNameSpace(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().list("5ktrial", "nh-sdk-ns", Context.NONE);
    }
}
