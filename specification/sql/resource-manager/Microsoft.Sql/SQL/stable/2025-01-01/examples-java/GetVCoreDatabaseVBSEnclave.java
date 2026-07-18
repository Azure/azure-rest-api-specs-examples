
/**
 * Samples for Databases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetVCoreDatabaseVBSEnclave.json
     */
    /**
     * Sample code: Gets a database configured with VBS enclave type.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsADatabaseConfiguredWithVBSEnclaveType(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().getWithResponse("Default-SQL-SouthEastAsia", "testsvr", "testdb", null,
            null, com.azure.core.util.Context.NONE);
    }
}
