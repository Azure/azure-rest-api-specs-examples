
import com.azure.resourcemanager.storage.models.StorageAccountRegenerateKeyParameters;

/**
 * Samples for StorageAccounts RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountRegenerateKerbKey.json
     */
    /**
     * Sample code: StorageAccountRegenerateKerbKey.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountRegenerateKerbKey(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().regenerateKeyWithResponse("res4167", "sto3539",
            new StorageAccountRegenerateKeyParameters().withKeyName("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
