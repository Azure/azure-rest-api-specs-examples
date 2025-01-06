
/**
 * Samples for Administrators Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * AdministratorDelete.json
     */
    /**
     * Sample code: AdministratorDelete.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        administratorDelete(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.administrators().delete("testrg", "testserver", "oooooooo-oooo-oooo-oooo-oooooooooooo",
            com.azure.core.util.Context.NONE);
    }
}
