
import com.azure.resourcemanager.sql.models.DatabaseExtensions;
import com.azure.resourcemanager.sql.models.OperationMode;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for DatabaseExtensionsOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateDatabaseExtensions.json
     */
    /**
     * Sample code: Create or Update database extensions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateDatabaseExtensions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseExtensionsOperations().createOrUpdate(
            "rg_20cbe0f0-c2d9-4522-9177-5469aad53029", "srv_1ffd1cf8-9951-47fb-807d-a9c384763849",
            "878e303f-1ea0-4f17-aa3d-a22ac5e9da08", "polybaseimport",
            new DatabaseExtensions().withOperationMode(OperationMode.POLYBASE_IMPORT)
                .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY).withStorageKey("fakeTokenPlaceholder")
                .withStorageUri("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml"),
            com.azure.core.util.Context.NONE);
    }
}
