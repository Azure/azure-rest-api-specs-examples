import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/CreateVCoreDatabaseByServiceObjective.json
     */
    /**
     * Sample code: Creates a VCore database by specifying service objective name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsAVCoreDatabaseBySpecifyingServiceObjectiveName(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .createOrUpdate(
                "Default-SQL-SouthEastAsia",
                "testsvr",
                "testdb",
                new DatabaseInner()
                    .withLocation("southeastasia")
                    .withSku(new Sku().withName("BC").withFamily("Gen4").withCapacity(2)),
                Context.NONE);
    }
}
