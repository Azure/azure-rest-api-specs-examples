
/**
 * Samples for ManagedInstanceAdvancedThreatProtectionSettings ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceAdvancedThreatProtectionSettingsListByInstance.json
     */
    /**
     * Sample code: List the managed instance's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listTheManagedInstanceSAdvancedThreatProtectionSettings(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAdvancedThreatProtectionSettings()
            .listByInstance("threatprotection-4799", "threatprotection-6440", com.azure.core.util.Context.NONE);
    }
}
