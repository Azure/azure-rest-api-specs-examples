
import com.azure.resourcemanager.mysqlflexibleserver.models.AdvancedThreatProtectionName;

/**
 * Samples for AdvancedThreatProtectionSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * AdvancedThreatProtectionSettingsGet.json
     */
    /**
     * Sample code: Get a server's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void getAServerSAdvancedThreatProtectionSettings(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.advancedThreatProtectionSettings().getWithResponse("threatprotection-6852", "threatprotection-2080",
            AdvancedThreatProtectionName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
