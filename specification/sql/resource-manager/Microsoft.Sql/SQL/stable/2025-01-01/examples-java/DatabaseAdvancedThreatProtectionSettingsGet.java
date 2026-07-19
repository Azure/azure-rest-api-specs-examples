
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;

/**
 * Samples for DatabaseAdvancedThreatProtectionSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseAdvancedThreatProtectionSettingsGet.json
     */
    /**
     * Sample code: Get a database's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getADatabaseSAdvancedThreatProtectionSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseAdvancedThreatProtectionSettings().getWithResponse("threatprotection-6852",
            "threatprotection-2080", "testdb", AdvancedThreatProtectionName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
