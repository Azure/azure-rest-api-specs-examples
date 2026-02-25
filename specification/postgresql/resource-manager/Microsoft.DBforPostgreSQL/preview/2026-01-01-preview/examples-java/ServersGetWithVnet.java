
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersGetWithVnet.json
     */
    /**
     * Sample code: Get information about an existing server that is integrated into a virtual network provided by
     * customer.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAnExistingServerThatIsIntegratedIntoAVirtualNetworkProvidedByCustomer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
