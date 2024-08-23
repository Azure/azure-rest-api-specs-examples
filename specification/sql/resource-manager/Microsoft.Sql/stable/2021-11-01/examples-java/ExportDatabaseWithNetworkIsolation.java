
import com.azure.resourcemanager.sql.models.ExportDatabaseDefinition;
import com.azure.resourcemanager.sql.models.NetworkIsolationSettings;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Databases Export.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ExportDatabaseWithNetworkIsolation.
     * json
     */
    /**
     * Sample code: Exports a database, using private link to communicate with SQL server and storage account.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportsADatabaseUsingPrivateLinkToCommunicateWithSQLServerAndStorageAccount(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases()
            .export("Default-SQL-SouthEastAsia", "testsvr", "testdb", new ExportDatabaseDefinition()
                .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY).withStorageKey("fakeTokenPlaceholder")
                .withStorageUri("https://test.blob.core.windows.net/test.bacpac").withAdministratorLogin("login")
                .withAdministratorLoginPassword("fakeTokenPlaceholder").withAuthenticationType("Sql")
                .withNetworkIsolation(new NetworkIsolationSettings().withStorageAccountResourceId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Storage/storageAccounts/test-privatelink")
                    .withSqlServerResourceId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr")),
                com.azure.core.util.Context.NONE);
    }
}
