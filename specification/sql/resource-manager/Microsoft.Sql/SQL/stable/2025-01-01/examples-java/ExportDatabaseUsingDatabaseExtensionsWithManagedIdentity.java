
import com.azure.resourcemanager.sql.models.DatabaseExtensions;
import com.azure.resourcemanager.sql.models.OperationMode;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/**
 * Samples for DatabaseExtensionsOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExportDatabaseUsingDatabaseExtensionsWithManagedIdentity.json
     */
    /**
     * Sample code: Export database using database extension with Managed Identity.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void exportDatabaseUsingDatabaseExtensionWithManagedIdentity(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseExtensionsOperations().createOrUpdate(
            "rg_d1ef9eae-044d-4710-ba59-b82e84ad3157", "srv_9243d320-ac4e-4f97-8e06-b1167dae5f4c",
            "db_7fe424c8-23cf-4ac3-bdc3-e21f424bdb68", "Export",
            new DatabaseExtensions().withOperationMode(OperationMode.EXPORT)
                .withStorageKeyType(StorageKeyType.MANAGED_IDENTITY).withStorageKey("fakeTokenPlaceholder")
                .withStorageUri("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml")
                .withAdministratorLogin(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName")
                .withAuthenticationType("ManagedIdentity"),
            com.azure.core.util.Context.NONE);
    }
}
