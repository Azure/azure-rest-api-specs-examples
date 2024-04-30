
/**
 * Samples for AdvancedThreatProtectionSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * AdvancedThreatProtectionSettingsList.json
     */
    /**
     * Sample code: Get list of server's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void getListOfServerSAdvancedThreatProtectionSettings(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.advancedThreatProtectionSettings().list("threatprotection-6852", "threatprotection-2080",
            com.azure.core.util.Context.NONE);
    }
}
