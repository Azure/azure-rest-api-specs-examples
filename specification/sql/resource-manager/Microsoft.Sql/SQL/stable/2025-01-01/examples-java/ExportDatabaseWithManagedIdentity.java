
import com.azure.resourcemanager.sql.models.ExportDatabaseDefinition;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for Databases Export.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExportDatabaseWithManagedIdentity.json
     */
    /**
     * Sample code: Exports a database, using Managed Identity to communicate with SQL server and storage account.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void exportsADatabaseUsingManagedIdentityToCommunicateWithSQLServerAndStorageAccount(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().export("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new ExportDatabaseDefinition().withStorageKeyType(StorageKeyType.MANAGED_IDENTITY)
                .withStorageKey("fakeTokenPlaceholder").withStorageUri("https://test.blob.core.windows.net/test.bacpac")
                .withAdministratorLogin(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName")
                .withAuthenticationType("ManagedIdentity"),
            com.azure.core.util.Context.NONE);
    }
}
