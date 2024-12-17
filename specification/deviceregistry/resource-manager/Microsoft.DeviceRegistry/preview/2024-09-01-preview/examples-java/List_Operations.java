
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_Operations.json
     */
    /**
     * Sample code: List_Operations.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listOperations(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
