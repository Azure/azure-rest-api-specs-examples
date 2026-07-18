
/**
 * Samples for ManagedDatabaseSecurityEvents ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseSecurityEventsGetMin.json
     */
    /**
     * Sample code: Get the managed database's security events with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getTheManagedDatabaseSSecurityEventsWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSecurityEvents().listByDatabase("testrg", "testcl", "database1", null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
