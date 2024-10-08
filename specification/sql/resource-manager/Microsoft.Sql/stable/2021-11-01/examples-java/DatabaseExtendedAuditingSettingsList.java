
/**
 * Samples for ExtendedDatabaseBlobAuditingPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseExtendedAuditingSettingsList.
     * json
     */
    /**
     * Sample code: List extended auditing settings of a database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExtendedAuditingSettingsOfADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getExtendedDatabaseBlobAuditingPolicies().listByDatabase(
            "blobauditingtest-6852", "blobauditingtest-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
