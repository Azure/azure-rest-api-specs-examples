/** Samples for Administrators ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/AdministratorsListByServer.json
     */
    /**
     * Sample code: AdministratorsListByServer.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void administratorsListByServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.administrators().listByServer("testrg", "pgtestsvc1", com.azure.core.util.Context.NONE);
    }
}
