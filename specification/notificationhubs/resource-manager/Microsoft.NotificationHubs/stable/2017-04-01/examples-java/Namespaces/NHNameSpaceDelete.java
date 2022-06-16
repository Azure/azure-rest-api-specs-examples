import com.azure.core.util.Context;

/** Samples for Namespaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceDelete.json
     */
    /**
     * Sample code: NameSpaceDelete.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceDelete(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().delete("5ktrial", "nh-sdk-ns", Context.NONE);
    }
}
