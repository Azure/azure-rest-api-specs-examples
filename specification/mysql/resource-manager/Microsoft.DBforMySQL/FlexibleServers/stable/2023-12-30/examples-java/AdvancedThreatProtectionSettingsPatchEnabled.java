
import com.azure.resourcemanager.mysqlflexibleserver.models.AdvancedThreatProtectionForUpdate;
import com.azure.resourcemanager.mysqlflexibleserver.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.mysqlflexibleserver.models.AdvancedThreatProtectionState;

/**
 * Samples for AdvancedThreatProtectionSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * AdvancedThreatProtectionSettingsPatchEnabled.json
     */
    /**
     * Sample code: Enable a server's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void enableAServerSAdvancedThreatProtectionSettings(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.advancedThreatProtectionSettings().update("threatprotection-4799", "threatprotection-6440",
            AdvancedThreatProtectionName.DEFAULT,
            new AdvancedThreatProtectionForUpdate().withState(AdvancedThreatProtectionState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
