
/**
 * Samples for Namespaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/ListByResourceGroup.json
     */
    /**
     * Sample code: Namespaces_List.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void namespacesList(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().listByResourceGroup("5ktrial", null, null, com.azure.core.util.Context.NONE);
    }
}
