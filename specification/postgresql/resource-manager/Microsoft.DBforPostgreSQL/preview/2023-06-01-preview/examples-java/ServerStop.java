
/**
 * Samples for Servers Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/
     * ServerStop.json
     */
    /**
     * Sample code: ServerStop.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverStop(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().stop("testrg", "testserver", com.azure.core.util.Context.NONE);
    }
}
