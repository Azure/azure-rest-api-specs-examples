/** Samples for NetworkConnections RunHealthChecks. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/NetworkConnections_RunHealthChecks.json
     */
    /**
     * Sample code: NetworkConnections_RunHealthChecks.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsRunHealthChecks(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.networkConnections().runHealthChecks("rg1", "uswest3network", com.azure.core.util.Context.NONE);
    }
}
