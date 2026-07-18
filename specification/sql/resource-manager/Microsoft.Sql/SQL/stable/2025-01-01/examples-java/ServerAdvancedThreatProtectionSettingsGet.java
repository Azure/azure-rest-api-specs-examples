
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;

/**
 * Samples for ServerAdvancedThreatProtectionSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAdvancedThreatProtectionSettingsGet.json
     */
    /**
     * Sample code: Get a server's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAServerSAdvancedThreatProtectionSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAdvancedThreatProtectionSettings().getWithResponse("threatprotection-4799",
            "threatprotection-6440", AdvancedThreatProtectionName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
