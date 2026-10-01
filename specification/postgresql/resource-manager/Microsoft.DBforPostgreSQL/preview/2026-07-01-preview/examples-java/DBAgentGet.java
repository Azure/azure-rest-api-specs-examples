
/**
 * Samples for DbAgents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/DBAgentGet.json
     */
    /**
     * Sample code: Get the singleton Default database agent configuration for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getTheSingletonDefaultDatabaseAgentConfigurationForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.dbAgents().getWithResponse("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
