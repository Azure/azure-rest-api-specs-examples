
/**
 * Samples for AdvancedThreatProtectionSettings ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/AdvancedThreatProtectionSettingsListByServer.json
     */
    /**
     * Sample code: List state of advanced threat protection settings for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listStateOfAdvancedThreatProtectionSettingsForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.advancedThreatProtectionSettings().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
