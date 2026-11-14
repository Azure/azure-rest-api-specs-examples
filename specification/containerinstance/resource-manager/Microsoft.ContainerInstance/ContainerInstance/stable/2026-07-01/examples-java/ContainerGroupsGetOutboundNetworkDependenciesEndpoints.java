
/**
 * Samples for ContainerGroups GetOutboundNetworkDependenciesEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupsGetOutboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: ContainerGroupsGetOutboundNetworkDependenciesEndpoints.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupsGetOutboundNetworkDependenciesEndpoints(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainerGroups().getOutboundNetworkDependenciesEndpointsWithResponse("demo",
            "demo-cg", com.azure.core.util.Context.NONE);
    }
}
