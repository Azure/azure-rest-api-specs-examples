
/**
 * Samples for Namespaces GetPnsCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/PnsCredentialsGet.json
     */
    /**
     * Sample code: Namespaces_GetPnsCredentials.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        namespacesGetPnsCredentials(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().getPnsCredentialsWithResponse("5ktrial", "nh-sdk-ns", com.azure.core.util.Context.NONE);
    }
}
