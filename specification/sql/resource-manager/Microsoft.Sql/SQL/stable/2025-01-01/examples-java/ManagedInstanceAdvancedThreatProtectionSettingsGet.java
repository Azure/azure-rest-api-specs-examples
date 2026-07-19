
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;

/**
 * Samples for ManagedInstanceAdvancedThreatProtectionSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceAdvancedThreatProtectionSettingsGet.json
     */
    /**
     * Sample code: Get a managed instance's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAManagedInstanceSAdvancedThreatProtectionSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAdvancedThreatProtectionSettings().getWithResponse(
            "threatprotection-4799", "threatprotection-6440", AdvancedThreatProtectionName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
