
import com.azure.resourcemanager.sql.models.ExportDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Databases Export.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ExportDatabase.json
     */
    /**
     * Sample code: Exports a database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportsADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().export("Default-SQL-SouthEastAsia", "testsvr",
            "testdb",
            new ExportDatabaseDefinition().withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                .withStorageKey("fakeTokenPlaceholder").withStorageUri("https://test.blob.core.windows.net/test.bacpac")
                .withAdministratorLogin("login").withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withAuthenticationType("Sql"),
            com.azure.core.util.Context.NONE);
    }
}
