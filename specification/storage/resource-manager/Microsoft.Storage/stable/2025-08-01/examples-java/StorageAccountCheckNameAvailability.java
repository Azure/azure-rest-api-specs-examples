
import com.azure.resourcemanager.storage.models.StorageAccountCheckNameAvailabilityParameters;

/**
 * Samples for StorageAccounts CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountCheckNameAvailability.json
     */
    /**
     * Sample code: StorageAccountCheckNameAvailability.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountCheckNameAvailability(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().checkNameAvailabilityWithResponse(
            new StorageAccountCheckNameAvailabilityParameters().withName("sto3363"), com.azure.core.util.Context.NONE);
    }
}
