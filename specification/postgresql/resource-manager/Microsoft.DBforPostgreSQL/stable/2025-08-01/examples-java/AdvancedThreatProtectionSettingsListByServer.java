
/**
 * Samples for AdvancedThreatProtectionSettings ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * AdvancedThreatProtectionSettingsListByServer.json
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
