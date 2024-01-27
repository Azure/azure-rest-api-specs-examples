
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ImportExistingDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/** Samples for Databases ImportMethod. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ImportDatabase.json
     */
    /**
     * Sample code: Imports to an existing empty database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importsToAnExistingEmptyDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().importMethod("Default-SQL-SouthEastAsia", "testsvr",
            "testdb",
            new ImportExistingDatabaseDefinition().withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                .withStorageKey("fakeTokenPlaceholder").withStorageUri("https://test.blob.core.windows.net/test.bacpac")
                .withAdministratorLogin("login").withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withAuthenticationType("Sql"),
            Context.NONE);
    }
}
