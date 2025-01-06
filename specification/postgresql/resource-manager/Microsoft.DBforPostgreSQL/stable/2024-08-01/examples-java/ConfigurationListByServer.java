
/**
 * Samples for Configurations ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * ConfigurationListByServer.json
     */
    /**
     * Sample code: ConfigurationList.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void configurationList(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.configurations().listByServer("testrg", "testserver", com.azure.core.util.Context.NONE);
    }
}
