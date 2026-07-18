
import com.azure.resourcemanager.sql.fluent.models.ServerAdvancedThreatProtectionInner;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionState;

/**
 * Samples for ServerAdvancedThreatProtectionSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAdvancedThreatProtectionSettingsCreateMax.json
     */
    /**
     * Sample code: Update a server's Advanced Threat Protection settings with all parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAServerSAdvancedThreatProtectionSettingsWithAllParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAdvancedThreatProtectionSettings().createOrUpdate("threatprotection-4799",
            "threatprotection-6440", AdvancedThreatProtectionName.DEFAULT,
            new ServerAdvancedThreatProtectionInner().withState(AdvancedThreatProtectionState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
