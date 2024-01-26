
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.DatabaseLicenseType;
import com.azure.resourcemanager.sql.models.DatabaseUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases Update. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/PatchVCoreDatabase.json
     */
    /**
     * Sample code: Updates a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatesADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().update("Default-SQL-SouthEastAsia", "testsvr",
            "testdb", new DatabaseUpdate().withSku(new Sku().withName("BC_Gen4_4")).withMaxSizeBytes(1073741824L)
                .withLicenseType(DatabaseLicenseType.LICENSE_INCLUDED),
            Context.NONE);
    }
}
