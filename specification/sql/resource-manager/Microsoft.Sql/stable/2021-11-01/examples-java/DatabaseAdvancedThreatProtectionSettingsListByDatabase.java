
import com.azure.core.util.Context;

/** Samples for DatabaseAdvancedThreatProtectionSettings ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * DatabaseAdvancedThreatProtectionSettingsListByDatabase.json
     */
    /**
     * Sample code: Lists the database's Advanced Threat Protection settings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listsTheDatabaseSAdvancedThreatProtectionSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseAdvancedThreatProtectionSettings()
            .listByDatabase("threatprotection-6852", "threatprotection-2080", "testdb", Context.NONE);
    }
}
