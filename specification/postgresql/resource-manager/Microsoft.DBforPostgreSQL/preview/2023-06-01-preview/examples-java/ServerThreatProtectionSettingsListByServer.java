/** Samples for ServerThreatProtectionSettings ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/ServerThreatProtectionSettingsListByServer.json
     */
    /**
     * Sample code: Get a server's Advanced Threat Protection settings.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getAServerSAdvancedThreatProtectionSettings(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .serverThreatProtectionSettings()
            .listByServer("threatprotection-6852", "threatprotection-2080", com.azure.core.util.Context.NONE);
    }
}
