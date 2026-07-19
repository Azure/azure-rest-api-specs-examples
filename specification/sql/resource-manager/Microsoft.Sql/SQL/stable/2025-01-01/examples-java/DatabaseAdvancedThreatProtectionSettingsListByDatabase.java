
/**
 * Samples for DatabaseAdvancedThreatProtectionSettings ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseAdvancedThreatProtectionSettingsListByDatabase.json
     */
    /**
     * Sample code: Lists the database's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listsTheDatabaseSAdvancedThreatProtectionSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseAdvancedThreatProtectionSettings().listByDatabase("threatprotection-6852",
            "threatprotection-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
