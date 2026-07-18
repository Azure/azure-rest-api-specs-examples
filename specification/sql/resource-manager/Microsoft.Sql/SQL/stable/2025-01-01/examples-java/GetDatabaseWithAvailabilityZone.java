
/**
 * Samples for Databases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetDatabaseWithAvailabilityZone.json
     */
    /**
     * Sample code: Gets a database with Availability zone specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsADatabaseWithAvailabilityZoneSpecified(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().getWithResponse("Default-SQL-SouthEastAsia", "testsvr", "testdb", null,
            null, com.azure.core.util.Context.NONE);
    }
}
