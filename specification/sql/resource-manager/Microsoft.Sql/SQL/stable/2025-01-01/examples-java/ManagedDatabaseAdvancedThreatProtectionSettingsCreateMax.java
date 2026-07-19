
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseAdvancedThreatProtectionInner;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionState;

/**
 * Samples for ManagedDatabaseAdvancedThreatProtectionSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseAdvancedThreatProtectionSettingsCreateMax.json
     */
    /**
     * Sample code: Update a managed database's Advanced Threat Protection settings with all parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAManagedDatabaseSAdvancedThreatProtectionSettingsWithAllParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseAdvancedThreatProtectionSettings().createOrUpdateWithResponse(
            "threatprotection-4799", "threatprotection-6440", "testdb", AdvancedThreatProtectionName.DEFAULT,
            new ManagedDatabaseAdvancedThreatProtectionInner().withState(AdvancedThreatProtectionState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
