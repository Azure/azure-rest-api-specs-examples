
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NHOperationsList.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void operationsList(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
