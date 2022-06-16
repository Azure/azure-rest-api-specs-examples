import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.StorageAccountRegenerateKeyParameters;

/** Samples for StorageAccounts RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/StorageAccountRegenerateKerbKey.json
     */
    /**
     * Sample code: StorageAccountRegenerateKerbKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountRegenerateKerbKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .regenerateKeyWithResponse(
                "res4167", "sto3539", new StorageAccountRegenerateKeyParameters().withKeyName("kerb1"), Context.NONE);
    }
}
