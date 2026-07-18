
import com.azure.resourcemanager.sql.models.ImportExistingDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Databases ImportMethod.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ImportDatabaseWithManagedIdentity.json
     */
    /**
     * Sample code: Imports to an existing empty database, using Managed Identity to communicate with SQL server and
     * storage account.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void importsToAnExistingEmptyDatabaseUsingManagedIdentityToCommunicateWithSQLServerAndStorageAccount(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().importMethod("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new ImportExistingDatabaseDefinition().withStorageKeyType(StorageKeyType.MANAGED_IDENTITY)
                .withStorageKey("fakeTokenPlaceholder").withStorageUri("https://test.blob.core.windows.net/test.bacpac")
                .withAdministratorLogin(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName")
                .withAuthenticationType("ManagedIdentity"),
            com.azure.core.util.Context.NONE);
    }
}
