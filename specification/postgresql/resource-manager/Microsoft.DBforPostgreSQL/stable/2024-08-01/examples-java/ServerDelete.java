
/**
 * Samples for Servers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/ServerDelete.json
     */
    /**
     * Sample code: ServerDelete.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverDelete(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().delete("testrg", "testserver", com.azure.core.util.Context.NONE);
    }
}
