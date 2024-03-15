
/**
 * Samples for Namespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/Delete.json
     */
    /**
     * Sample code: Namespaces_Delete.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void namespacesDelete(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().deleteByResourceGroupWithResponse("5ktrial", "nh-sdk-ns",
            com.azure.core.util.Context.NONE);
    }
}
