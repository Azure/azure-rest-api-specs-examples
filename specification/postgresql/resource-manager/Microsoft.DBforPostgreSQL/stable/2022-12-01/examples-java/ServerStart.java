/** Samples for Servers Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/ServerStart.json
     */
    /**
     * Sample code: ServerStart.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverStart(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().start("testrg", "testserver", com.azure.core.util.Context.NONE);
    }
}
