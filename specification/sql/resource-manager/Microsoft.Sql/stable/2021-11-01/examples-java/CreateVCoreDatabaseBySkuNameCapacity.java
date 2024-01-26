
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateVCoreDatabaseBySkuNameCapacity.
     * json
     */
    /**
     * Sample code: Creates a VCore database by specifying sku name and capacity.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createsAVCoreDatabaseBySpecifyingSkuNameAndCapacity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdb",
            new DatabaseInner().withLocation("southeastasia").withSku(new Sku().withName("BC_Gen4").withCapacity(2)),
            Context.NONE);
    }
}
