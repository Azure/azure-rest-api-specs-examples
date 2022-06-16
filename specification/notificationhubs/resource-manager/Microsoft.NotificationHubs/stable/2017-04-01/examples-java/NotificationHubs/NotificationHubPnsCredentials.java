import com.azure.core.util.Context;

/** Samples for NotificationHubs GetPnsCredentials. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubPnsCredentials.json
     */
    /**
     * Sample code: notificationHubPnsCredentials.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubPnsCredentials(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().getPnsCredentialsWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub", Context.NONE);
    }
}
