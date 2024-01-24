
/**
 * Samples for DedicatedHsm ListOutboundNetworkDependenciesEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/
     * examples/GetOutboundNetworkDependenciesEndpointsList.json
     */
    /**
     * Sample code: List OutboundNetworkDependenciesEndpoints by Managed Cluster.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void listOutboundNetworkDependenciesEndpointsByManagedCluster(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().listOutboundNetworkDependenciesEndpoints("hsm-group", "hsm1",
            com.azure.core.util.Context.NONE);
    }
}
