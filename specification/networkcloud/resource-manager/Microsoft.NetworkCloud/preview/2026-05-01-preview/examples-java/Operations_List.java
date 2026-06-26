
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Operations_List.json
     */
    /**
     * Sample code: List resource provider operations.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listResourceProviderOperations(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
