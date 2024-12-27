
/**
 * Samples for Namespaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/
     * NHNameSpaceGet.json
     */
    /**
     * Sample code: NameSpaceGet.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceGet(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().getByResourceGroupWithResponse("5ktrial", "nh-sdk-ns", com.azure.core.util.Context.NONE);
    }
}
