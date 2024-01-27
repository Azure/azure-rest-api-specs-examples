
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ImportNewDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/** Samples for Servers ImportDatabase. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ImportNewDatabase.json
     */
    /**
     * Sample code: Imports to a new database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importsToANewDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().importDatabase("Default-SQL-SouthEastAsia", "testsvr",
            new ImportNewDatabaseDefinition().withDatabaseName("testdb")
                .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY).withStorageKey("fakeTokenPlaceholder")
                .withStorageUri("https://test.blob.core.windows.net/test.bacpac").withAdministratorLogin("login")
                .withAdministratorLoginPassword("fakeTokenPlaceholder").withAuthenticationType("Sql"),
            Context.NONE);
    }
}
