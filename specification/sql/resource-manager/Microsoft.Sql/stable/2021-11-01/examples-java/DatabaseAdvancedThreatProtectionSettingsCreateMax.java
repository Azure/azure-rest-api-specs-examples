
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseAdvancedThreatProtectionInner;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionName;
import com.azure.resourcemanager.sql.models.AdvancedThreatProtectionState;

/** Samples for DatabaseAdvancedThreatProtectionSettings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * DatabaseAdvancedThreatProtectionSettingsCreateMax.json
     */
    /**
     * Sample code: Update a database's Advanced Threat Protection settings with all parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateADatabaseSAdvancedThreatProtectionSettingsWithAllParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseAdvancedThreatProtectionSettings()
            .createOrUpdateWithResponse("threatprotection-4799", "threatprotection-6440", "testdb",
                AdvancedThreatProtectionName.DEFAULT,
                new DatabaseAdvancedThreatProtectionInner().withState(AdvancedThreatProtectionState.ENABLED),
                Context.NONE);
    }
}
