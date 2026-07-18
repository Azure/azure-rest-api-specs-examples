
/**
 * Samples for ManagedDatabaseAdvancedThreatProtectionSettings ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseAdvancedThreatProtectionSettingsListByDatabase.json
     */
    /**
     * Sample code: Get a list of the managed database's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAListOfTheManagedDatabaseSAdvancedThreatProtectionSettings(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseAdvancedThreatProtectionSettings().listByDatabase(
            "threatprotection-6852", "threatprotection-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
