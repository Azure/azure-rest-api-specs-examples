
/**
 * Samples for ServerAdvancedThreatProtectionSettings ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAdvancedThreatProtectionSettingsListByServer.json
     */
    /**
     * Sample code: List the server's Advanced Threat Protection settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listTheServerSAdvancedThreatProtectionSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAdvancedThreatProtectionSettings().listByServer("threatprotection-4799",
            "threatprotection-6440", com.azure.core.util.Context.NONE);
    }
}
