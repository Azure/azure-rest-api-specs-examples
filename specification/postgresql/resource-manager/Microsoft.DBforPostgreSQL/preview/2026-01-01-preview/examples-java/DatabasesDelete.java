
/**
 * Samples for Databases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/DatabasesDelete.json
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
