
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersGetWithPrivateEndpoints.json
     */
    /**
     * Sample code: Get information about an existing server that isn't integrated into a virtual network provided by
     * customer and has private endpoint connections.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getInformationAboutAnExistingServerThatIsnTIntegratedIntoAVirtualNetworkProvidedByCustomerAndHasPrivateEndpointConnections(
            com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
