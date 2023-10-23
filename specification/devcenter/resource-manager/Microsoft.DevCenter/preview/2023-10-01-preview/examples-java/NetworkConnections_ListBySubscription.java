/** Samples for NetworkConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/NetworkConnections_ListBySubscription.json
     */
    /**
     * Sample code: NetworkConnections_ListBySubscription.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsListBySubscription(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.networkConnections().list(null, com.azure.core.util.Context.NONE);
    }
}
