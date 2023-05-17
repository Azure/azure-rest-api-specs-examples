/** Samples for NetworkConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/NetworkConnections_ListBySubscription.json
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
