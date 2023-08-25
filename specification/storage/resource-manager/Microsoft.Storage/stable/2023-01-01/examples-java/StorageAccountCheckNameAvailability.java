import com.azure.resourcemanager.storage.models.StorageAccountCheckNameAvailabilityParameters;

/** Samples for StorageAccounts CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountCheckNameAvailability.json
     */
    /**
     * Sample code: StorageAccountCheckNameAvailability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountCheckNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .checkNameAvailabilityWithResponse(
                new StorageAccountCheckNameAvailabilityParameters().withName("sto3363"),
                com.azure.core.util.Context.NONE);
    }
}
