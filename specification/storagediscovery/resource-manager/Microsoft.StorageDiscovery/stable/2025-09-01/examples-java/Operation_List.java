
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Operation_List.json
     */
    /**
     * Sample code: List all provider operations.
     * 
     * @param manager Entry point to StorageDiscoveryManager.
     */
    public static void
        listAllProviderOperations(com.azure.resourcemanager.storagediscovery.StorageDiscoveryManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
