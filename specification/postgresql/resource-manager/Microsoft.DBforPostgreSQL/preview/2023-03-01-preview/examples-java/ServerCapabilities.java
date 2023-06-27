/** Samples for ServerCapabilities List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/ServerCapabilities.json
     */
    /**
     * Sample code: ServerCapabilitiesList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverCapabilitiesList(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.serverCapabilities().list("testrg", "pgtestsvc4", com.azure.core.util.Context.NONE);
    }
}
