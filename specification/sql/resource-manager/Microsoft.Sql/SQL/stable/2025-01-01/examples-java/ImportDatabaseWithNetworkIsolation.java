
import com.azure.resourcemanager.sql.models.ImportExistingDatabaseDefinition;
import com.azure.resourcemanager.sql.models.NetworkIsolationSettings;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Databases ImportMethod.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ImportDatabaseWithNetworkIsolation.json
     */
    /**
     * Sample code: Imports to an existing empty database, using private link to communicate with SQL server and storage
     * account.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void importsToAnExistingEmptyDatabaseUsingPrivateLinkToCommunicateWithSQLServerAndStorageAccount(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().importMethod("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new ImportExistingDatabaseDefinition().withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                .withStorageKey("fakeTokenPlaceholder").withStorageUri("https://test.blob.core.windows.net/test.bacpac")
                .withAdministratorLogin("login").withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withAuthenticationType("Sql")
                .withNetworkIsolation(new NetworkIsolationSettings().withStorageAccountResourceId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Storage/storageAccounts/test-privatelink")
                    .withSqlServerResourceId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr")),
            com.azure.core.util.Context.NONE);
    }
}
