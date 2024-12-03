
/**
 * Samples for NotificationHubs GetPnsCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/PnsCredentialsGet.json
     */
    /**
     * Sample code: NotificationHubs_GetPnsCredentials.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        notificationHubsGetPnsCredentials(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().getPnsCredentialsWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub",
            com.azure.core.util.Context.NONE);
    }
}
