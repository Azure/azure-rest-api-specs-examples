
import com.azure.resourcemanager.sql.fluent.models.ServerAdvancedThreatProtectionInner;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionState;

/**
 * Samples for ServerAdvancedThreatProtectionSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ServerAdvancedThreatProtectionSettingsCreateMin.json
     */
    /**
     * Sample code: Update a server's Advanced Threat Protection settings with minimal parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAServerSAdvancedThreatProtectionSettingsWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAdvancedThreatProtectionSettings().createOrUpdate(
            "threatprotection-4799", "threatprotection-6440", AdvancedThreatProtectionName.DEFAULT,
            new ServerAdvancedThreatProtectionInner().withState(AdvancedThreatProtectionState.DISABLED),
            com.azure.core.util.Context.NONE);
    }
}
