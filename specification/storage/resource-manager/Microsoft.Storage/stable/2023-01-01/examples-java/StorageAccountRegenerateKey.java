
import com.azure.resourcemanager.storage.models.StorageAccountRegenerateKeyParameters;

/** Samples for StorageAccounts RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountRegenerateKey.
     * json
     */
    /**
     * Sample code: StorageAccountRegenerateKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().regenerateKeyWithResponse("res4167",
            "sto3539", new StorageAccountRegenerateKeyParameters().withKeyName("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
