
import com.azure.resourcemanager.notificationhubs.models.CheckAvailabilityParameters;

/**
 * Samples for NotificationHubs CheckNotificationHubAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/
     * NotificationHubs/NotificationHubCheckNameAvailability.json
     */
    /**
     * Sample code: notificationHubCheckNameAvailability.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubCheckNameAvailability(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().checkNotificationHubAvailabilityWithResponse("5ktrial", "locp-newns",
            new CheckAvailabilityParameters().withName("sdktest").withLocation("West Europe"),
            com.azure.core.util.Context.NONE);
    }
}
