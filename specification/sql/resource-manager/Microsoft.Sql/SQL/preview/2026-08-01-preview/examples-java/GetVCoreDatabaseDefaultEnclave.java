
/**
 * Samples for Databases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/GetVCoreDatabaseDefaultEnclave.json
     */
    /**
     * Sample code: Gets a database configured with Default enclave type.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsADatabaseConfiguredWithDefaultEnclaveType(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().getWithResponse("Default-SQL-SouthEastAsia", "testsvr", "testdb", null,
            null, com.azure.core.util.Context.NONE);
    }
}
