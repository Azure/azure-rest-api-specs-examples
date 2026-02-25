
import com.azure.resourcemanager.postgresqlflexibleserver.models.ThreatProtectionName;

/**
 * Samples for AdvancedThreatProtectionSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/AdvancedThreatProtectionSettingsGet.json
     */
    /**
     * Sample code: Get state of advanced threat protection settings for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getStateOfAdvancedThreatProtectionSettingsForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.advancedThreatProtectionSettings().getWithResponse("exampleresourcegroup", "exampleserver",
            ThreatProtectionName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
