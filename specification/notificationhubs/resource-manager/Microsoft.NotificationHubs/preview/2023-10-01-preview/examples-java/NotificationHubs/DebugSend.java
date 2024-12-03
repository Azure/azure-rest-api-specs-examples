
/**
 * Samples for NotificationHubs DebugSend.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/DebugSend.json
     */
    /**
     * Sample code: NotificationHubs_DebugSend.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        notificationHubsDebugSend(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().debugSendWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub",
            com.azure.core.util.Context.NONE);
    }
}
