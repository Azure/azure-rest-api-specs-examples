
import com.azure.resourcemanager.sql.models.DatabaseExtensions;
import com.azure.resourcemanager.sql.models.OperationMode;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for DatabaseExtensionsOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ExportDatabaseUsingDatabaseExtensions
     * .json
     */
    /**
     * Sample code: Export database using database extension.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportDatabaseUsingDatabaseExtension(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseExtensionsOperations().createOrUpdate(
            "rg_d1ef9eae-044d-4710-ba59-b82e84ad3157", "srv_9243d320-ac4e-4f97-8e06-b1167dae5f4c",
            "db_7fe424c8-23cf-4ac3-bdc3-e21f424bdb68", "Export",
            new DatabaseExtensions().withOperationMode(OperationMode.EXPORT)
                .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY).withStorageKey("fakeTokenPlaceholder")
                .withStorageUri("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml")
                .withAdministratorLogin("login").withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withAuthenticationType("Sql"),
            com.azure.core.util.Context.NONE);
    }
}
