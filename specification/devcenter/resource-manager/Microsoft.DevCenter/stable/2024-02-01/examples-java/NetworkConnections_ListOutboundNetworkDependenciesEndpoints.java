
/**
 * Samples for NetworkConnections ListOutboundNetworkDependenciesEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * NetworkConnections_ListOutboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: ListOutboundNetworkDependencies.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void listOutboundNetworkDependencies(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.networkConnections().listOutboundNetworkDependenciesEndpoints("rg1", "uswest3network", null,
            com.azure.core.util.Context.NONE);
    }
}
