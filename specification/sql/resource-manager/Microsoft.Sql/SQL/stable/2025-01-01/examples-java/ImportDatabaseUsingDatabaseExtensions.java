
import com.azure.resourcemanager.sql.models.DatabaseExtensions;
import com.azure.resourcemanager.sql.models.OperationMode;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for DatabaseExtensionsOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ImportDatabaseUsingDatabaseExtensions.json
     */
    /**
     * Sample code: Import database using database extension.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void importDatabaseUsingDatabaseExtension(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseExtensionsOperations().createOrUpdate(
            "rg_062866bf-c4f4-41f9-abf0-b59132ca7924", "srv_2d6be2d2-26c8-4930-8fb6-82a5e95e0e82",
            "db_2a47e946-e414-4c00-94ac-ed886bb78302", "Import",
            new DatabaseExtensions().withOperationMode(OperationMode.IMPORT)
                .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY).withStorageKey("fakeTokenPlaceholder")
                .withStorageUri("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml")
                .withAdministratorLogin("login").withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withAuthenticationType("Sql"),
            com.azure.core.util.Context.NONE);
    }
}
