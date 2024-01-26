
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * CreateDatabaseMaintenanceConfiguration.json
     */
    /**
     * Sample code: Creates a database with preferred maintenance window.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createsADatabaseWithPreferredMaintenanceWindow(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdb",
            new DatabaseInner().withLocation("southeastasia").withSku(new Sku().withName("S2").withTier("Standard"))
                .withCreateMode(CreateMode.DEFAULT).withCollation("SQL_Latin1_General_CP1_CI_AS")
                .withMaxSizeBytes(1073741824L).withMaintenanceConfigurationId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_SouthEastAsia_1"),
            Context.NONE);
    }
}
