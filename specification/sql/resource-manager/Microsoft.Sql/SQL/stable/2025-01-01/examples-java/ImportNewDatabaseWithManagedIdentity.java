
import com.azure.resourcemanager.sql.models.ImportNewDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Servers ImportDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ImportNewDatabaseWithManagedIdentity.json
     */
    /**
     * Sample code: Imports to a new database, using Managed Identity for the SQL server and storage account.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void importsToANewDatabaseUsingManagedIdentityForTheSQLServerAndStorageAccount(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().importDatabase("Default-SQL-SouthEastAsia", "testsvr",
            new ImportNewDatabaseDefinition().withDatabaseName("testdb")
                .withStorageKeyType(StorageKeyType.MANAGED_IDENTITY).withStorageKey("fakeTokenPlaceholder")
                .withStorageUri("https://test.blob.core.windows.net/test.bacpac")
                .withAdministratorLogin(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName")
                .withAuthenticationType("ManagedIdentity"),
            com.azure.core.util.Context.NONE);
    }
}
