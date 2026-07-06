
/**
 * Samples for PrivateLinkResources ListByStorageAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountListPrivateLinkResources.json
     */
    /**
     * Sample code: StorageAccountListPrivateLinkResources.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountListPrivateLinkResources(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getPrivateLinkResources().listByStorageAccountWithResponse("res6977", "sto2527",
            com.azure.core.util.Context.NONE);
    }
}
