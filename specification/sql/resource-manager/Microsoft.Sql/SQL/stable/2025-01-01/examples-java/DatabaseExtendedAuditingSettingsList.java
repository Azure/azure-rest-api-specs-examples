
/**
 * Samples for ExtendedDatabaseBlobAuditingPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseExtendedAuditingSettingsList.json
     */
    /**
     * Sample code: List extended auditing settings of a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listExtendedAuditingSettingsOfADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getExtendedDatabaseBlobAuditingPolicies().listByDatabase("blobauditingtest-6852",
            "blobauditingtest-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
