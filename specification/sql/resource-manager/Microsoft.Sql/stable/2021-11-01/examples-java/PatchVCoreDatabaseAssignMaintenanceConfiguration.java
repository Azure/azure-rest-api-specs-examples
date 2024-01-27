
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.DatabaseUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * PatchVCoreDatabaseAssignMaintenanceConfiguration.json
     */
    /**
     * Sample code: Assigns maintenance window to a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void assignsMaintenanceWindowToADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().update("Default-SQL-SouthEastAsia", "testsvr",
            "testdb",
            new DatabaseUpdate().withSku(new Sku().withName("BC_Gen5_4")).withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_SouthEastAsia_1"),
            Context.NONE);
    }
}
