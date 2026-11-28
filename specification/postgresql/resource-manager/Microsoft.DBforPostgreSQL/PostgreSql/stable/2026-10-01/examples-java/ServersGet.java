
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ServersGet.json
     */
    /**
     * Sample code: Get information about an existing server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAnExistingServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
