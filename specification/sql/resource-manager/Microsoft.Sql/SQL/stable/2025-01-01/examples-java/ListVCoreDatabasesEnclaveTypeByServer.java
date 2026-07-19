
/**
 * Samples for Databases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListVCoreDatabasesEnclaveTypeByServer.json
     */
    /**
     * Sample code: Gets a list of databases configured with enclave type.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsAListOfDatabasesConfiguredWithEnclaveType(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().listByServer("Default-SQL-SouthEastAsia", "testsvr", null,
            com.azure.core.util.Context.NONE);
    }
}
