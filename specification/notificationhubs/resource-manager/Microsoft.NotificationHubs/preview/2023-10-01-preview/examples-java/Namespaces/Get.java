
/**
 * Samples for Namespaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/Get.json
     */
    /**
     * Sample code: Namespaces_Get.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void namespacesGet(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().getByResourceGroupWithResponse("5ktrial", "nh-sdk-ns", com.azure.core.util.Context.NONE);
    }
}
