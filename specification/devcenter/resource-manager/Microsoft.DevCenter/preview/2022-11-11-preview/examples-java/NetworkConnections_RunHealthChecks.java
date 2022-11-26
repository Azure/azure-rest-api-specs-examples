import com.azure.core.util.Context;

/** Samples for NetworkConnections RunHealthChecks. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/NetworkConnections_RunHealthChecks.json
     */
    /**
     * Sample code: NetworkConnections_RunHealthChecks.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsRunHealthChecks(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.networkConnections().runHealthChecks("rg1", "uswest3network", Context.NONE);
    }
}
