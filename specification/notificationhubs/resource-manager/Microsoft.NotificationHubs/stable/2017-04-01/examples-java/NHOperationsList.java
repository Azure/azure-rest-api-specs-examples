import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NHOperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void operationsList(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.operations().list(Context.NONE);
    }
}
