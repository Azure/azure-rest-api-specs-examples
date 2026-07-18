
/**
 * Samples for DatabaseBlobAuditingPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseAuditingSettingsList.json
     */
    /**
     * Sample code: List audit settings of a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listAuditSettingsOfADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseBlobAuditingPolicies().listByDatabase("blobauditingtest-6852",
            "blobauditingtest-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
