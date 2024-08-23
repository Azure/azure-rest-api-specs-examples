
/**
 * Samples for ServerAdvancedThreatProtectionSettings ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ServerAdvancedThreatProtectionSettingsListByServer.json
     */
    /**
     * Sample code: List the server's Advanced Threat Protection settings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listTheServerSAdvancedThreatProtectionSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAdvancedThreatProtectionSettings()
            .listByServer("threatprotection-4799", "threatprotection-6440", com.azure.core.util.Context.NONE);
    }
}
