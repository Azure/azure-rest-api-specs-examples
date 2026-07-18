
import com.azure.resourcemanager.sql.models.ImportNewDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Servers ImportDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ImportNewDatabase.json
     */
    /**
     * Sample code: Imports to a new database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void importsToANewDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().importDatabase("Default-SQL-SouthEastAsia", "testsvr",
            new ImportNewDatabaseDefinition().withDatabaseName("testdb")
                .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY).withStorageKey("fakeTokenPlaceholder")
                .withStorageUri("https://test.blob.core.windows.net/test.bacpac").withAdministratorLogin("login")
                .withAdministratorLoginPassword("fakeTokenPlaceholder").withAuthenticationType("Sql"),
            com.azure.core.util.Context.NONE);
    }
}
