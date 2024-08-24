
import com.azure.resourcemanager.sql.models.DatabaseUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for Databases Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * PatchVCoreDatabaseResetMaintenanceConfiguration.json
     */
    /**
     * Sample code: Resets maintenance window of a database to default.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        resetsMaintenanceWindowOfADatabaseToDefault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().update("Default-SQL-SouthEastAsia", "testsvr",
            "testdb",
            new DatabaseUpdate().withSku(new Sku().withName("BC_Gen5_4")).withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
            com.azure.core.util.Context.NONE);
    }
}
