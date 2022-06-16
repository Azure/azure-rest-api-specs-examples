import com.azure.core.util.Context;

/** Samples for Namespaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceList.json
     */
    /**
     * Sample code: NameSpaceList.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceList(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().list(Context.NONE);
    }
}
