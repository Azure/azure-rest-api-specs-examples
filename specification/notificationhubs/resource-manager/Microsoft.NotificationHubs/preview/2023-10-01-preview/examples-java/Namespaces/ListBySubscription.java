
/**
 * Samples for Namespaces List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/ListBySubscription.json
     */
    /**
     * Sample code: Namespaces_ListAll.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void namespacesListAll(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().list(null, null, com.azure.core.util.Context.NONE);
    }
}
