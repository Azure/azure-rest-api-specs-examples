
/**
 * Samples for Configurations ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * ConfigurationsListByServer.json
     */
    /**
     * Sample code: List all configurations (also known as server parameters) of a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAllConfigurationsAlsoKnownAsServerParametersOfAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.configurations().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
