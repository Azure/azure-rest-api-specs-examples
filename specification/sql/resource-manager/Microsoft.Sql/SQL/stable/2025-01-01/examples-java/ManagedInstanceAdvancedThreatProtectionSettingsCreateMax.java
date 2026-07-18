
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceAdvancedThreatProtectionInner;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionState;

/**
 * Samples for ManagedInstanceAdvancedThreatProtectionSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceAdvancedThreatProtectionSettingsCreateMax.json
     */
    /**
     * Sample code: Update a managed instance's Advanced Threat Protection settings with all parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAManagedInstanceSAdvancedThreatProtectionSettingsWithAllParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAdvancedThreatProtectionSettings().createOrUpdate(
            "threatprotection-4799", "threatprotection-6440", AdvancedThreatProtectionName.DEFAULT,
            new ManagedInstanceAdvancedThreatProtectionInner().withState(AdvancedThreatProtectionState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
