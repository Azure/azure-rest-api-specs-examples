import com.azure.core.util.Context;

/** Samples for NetworkConnections ListHealthDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/NetworkConnections_ListHealthDetails.json
     */
    /**
     * Sample code: NetworkConnections_ListHealthDetails.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsListHealthDetails(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.networkConnections().listHealthDetails("rg1", "uswest3network", null, Context.NONE);
    }
}
