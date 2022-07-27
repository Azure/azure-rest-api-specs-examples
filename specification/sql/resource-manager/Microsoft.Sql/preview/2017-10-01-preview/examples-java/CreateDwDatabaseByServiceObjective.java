import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/CreateDwDatabaseByServiceObjective.json
     */
    /**
     * Sample code: Creates a data warehouse by specifying service objective name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADataWarehouseBySpecifyingServiceObjectiveName(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .createOrUpdate(
                "Default-SQL-SouthEastAsia",
                "testsvr",
                "testdw",
                new DatabaseInner().withLocation("westus").withSku(new Sku().withName("DW1000c")),
                Context.NONE);
    }
}
