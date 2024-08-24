
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;

/**
 * Samples for DatabaseAdvancedThreatProtectionSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * DatabaseAdvancedThreatProtectionSettingsGet.json
     */
    /**
     * Sample code: Get a database's Advanced Threat Protection settings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getADatabaseSAdvancedThreatProtectionSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseAdvancedThreatProtectionSettings().getWithResponse(
            "threatprotection-6852", "threatprotection-2080", "testdb", AdvancedThreatProtectionName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
