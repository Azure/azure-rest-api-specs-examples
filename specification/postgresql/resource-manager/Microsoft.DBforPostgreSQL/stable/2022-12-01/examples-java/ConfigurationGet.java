/** Samples for Configurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/ConfigurationGet.json
     */
    /**
     * Sample code: ConfigurationGet.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void configurationGet(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .configurations()
            .getWithResponse("testrg", "testserver", "array_nulls", com.azure.core.util.Context.NONE);
    }
}
