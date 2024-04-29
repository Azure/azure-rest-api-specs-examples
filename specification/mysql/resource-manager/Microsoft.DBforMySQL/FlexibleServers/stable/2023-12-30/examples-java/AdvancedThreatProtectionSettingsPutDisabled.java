
import com.azure.resourcemanager.mysqlflexibleserver.fluent.models.AdvancedThreatProtectionInner;
import com.azure.resourcemanager.mysqlflexibleserver.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.mysqlflexibleserver.models.AdvancedThreatProtectionState;

/**
 * Samples for AdvancedThreatProtectionSettings UpdatePut.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * AdvancedThreatProtectionSettingsPutDisabled.json
     */
    /**
     * Sample code: Disable a server's Advanced Threat Protection settings with all parameters using PUT.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void disableAServerSAdvancedThreatProtectionSettingsWithAllParametersUsingPUT(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.advancedThreatProtectionSettings().updatePut("threatprotection-4799", "threatprotection-6440",
            AdvancedThreatProtectionName.DEFAULT,
            new AdvancedThreatProtectionInner().withState(AdvancedThreatProtectionState.DISABLED),
            com.azure.core.util.Context.NONE);
    }
}
