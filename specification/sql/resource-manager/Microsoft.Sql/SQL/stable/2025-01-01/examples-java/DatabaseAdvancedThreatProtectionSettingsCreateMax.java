
import com.azure.resourcemanager.sql.fluent.models.DatabaseAdvancedThreatProtectionInner;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionState;

/**
 * Samples for DatabaseAdvancedThreatProtectionSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseAdvancedThreatProtectionSettingsCreateMax.json
     */
    /**
     * Sample code: Update a database's Advanced Threat Protection settings with all parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateADatabaseSAdvancedThreatProtectionSettingsWithAllParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseAdvancedThreatProtectionSettings().createOrUpdateWithResponse(
            "threatprotection-4799", "threatprotection-6440", "testdb", AdvancedThreatProtectionName.DEFAULT,
            new DatabaseAdvancedThreatProtectionInner().withState(AdvancedThreatProtectionState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
