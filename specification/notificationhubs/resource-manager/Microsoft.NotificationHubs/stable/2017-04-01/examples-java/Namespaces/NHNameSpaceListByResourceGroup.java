import com.azure.core.util.Context;

/** Samples for Namespaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceListByResourceGroup.json
     */
    /**
     * Sample code: NameSpaceListByResourceGroup.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceListByResourceGroup(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().listByResourceGroup("5ktrial", Context.NONE);
    }
}
