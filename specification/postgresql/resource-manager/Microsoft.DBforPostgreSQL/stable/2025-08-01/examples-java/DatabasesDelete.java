
/**
 * Samples for Databases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/DatabasesDelete.
     * json
     */
    /**
     * Sample code: Delete an existing database.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        deleteAnExistingDatabase(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.databases().delete("exampleresourcegroup", "exampleserver", "exampledatabase",
            com.azure.core.util.Context.NONE);
    }
}
