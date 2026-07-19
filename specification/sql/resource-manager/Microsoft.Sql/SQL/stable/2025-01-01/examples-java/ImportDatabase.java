
import com.azure.resourcemanager.sql.models.ImportExistingDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Databases ImportMethod.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ImportDatabase.json
     */
    /**
     * Sample code: Imports to an existing empty database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void importsToAnExistingEmptyDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().importMethod("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new ImportExistingDatabaseDefinition().withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                .withStorageKey("fakeTokenPlaceholder").withStorageUri("https://test.blob.core.windows.net/test.bacpac")
                .withAdministratorLogin("login").withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withAuthenticationType("Sql"),
            com.azure.core.util.Context.NONE);
    }
}
