
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;

/**
 * Samples for ManagedDatabaseAdvancedThreatProtectionSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseAdvancedThreatProtectionSettingsGet.json
     */
    /**
     * Sample code: Get a managed database's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAManagedDatabaseSAdvancedThreatProtectionSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseAdvancedThreatProtectionSettings().getWithResponse(
            "threatprotection-6852", "threatprotection-2080", "testdb", AdvancedThreatProtectionName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
