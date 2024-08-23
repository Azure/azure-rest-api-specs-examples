
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;

/**
 * Samples for ServerAdvancedThreatProtectionSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ServerAdvancedThreatProtectionSettingsGet.json
     */
    /**
     * Sample code: Get a server's Advanced Threat Protection settings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAServerSAdvancedThreatProtectionSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAdvancedThreatProtectionSettings().getWithResponse(
            "threatprotection-4799", "threatprotection-6440", AdvancedThreatProtectionName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
