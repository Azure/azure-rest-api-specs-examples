
import com.azure.resourcemanager.sql.models.ExportDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Databases Export.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExportDatabase.json
     */
    /**
     * Sample code: Exports a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void exportsADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().export("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new ExportDatabaseDefinition().withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                .withStorageKey("fakeTokenPlaceholder").withStorageUri("https://test.blob.core.windows.net/test.bacpac")
                .withAdministratorLogin("login").withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withAuthenticationType("Sql"),
            com.azure.core.util.Context.NONE);
    }
}
