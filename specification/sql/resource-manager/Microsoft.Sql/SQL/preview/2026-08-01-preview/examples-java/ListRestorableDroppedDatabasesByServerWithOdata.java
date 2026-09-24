
/**
 * Samples for RestorableDroppedDatabases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ListRestorableDroppedDatabasesByServerWithOdata.json
     */
    /**
     * Sample code: Gets a list of restorable dropped databases with OData filtering.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfRestorableDroppedDatabasesWithODataFiltering(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorableDroppedDatabases().listByServer("Default-SQL-SouthEastAsia", "testsvr",
            "eyJuYW1lIjoidGVzdERiMCwxMzE1OTIzODQwMDAwMDAwMDAifQ==", 25L, com.azure.core.util.Context.NONE);
    }
}
