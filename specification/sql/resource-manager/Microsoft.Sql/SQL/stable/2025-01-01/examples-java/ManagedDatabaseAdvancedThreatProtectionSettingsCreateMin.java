
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseAdvancedThreatProtectionInner;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionState;

/**
 * Samples for ManagedDatabaseAdvancedThreatProtectionSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseAdvancedThreatProtectionSettingsCreateMin.json
     */
    /**
     * Sample code: Update a managed database's Advanced Threat Protection settings with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAManagedDatabaseSAdvancedThreatProtectionSettingsWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseAdvancedThreatProtectionSettings().createOrUpdateWithResponse(
            "threatprotection-4799", "threatprotection-6440", "testdb", AdvancedThreatProtectionName.DEFAULT,
            new ManagedDatabaseAdvancedThreatProtectionInner().withState(AdvancedThreatProtectionState.DISABLED),
            com.azure.core.util.Context.NONE);
    }
}
