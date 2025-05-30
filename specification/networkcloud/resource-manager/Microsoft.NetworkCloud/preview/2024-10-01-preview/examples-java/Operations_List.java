
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * Operations_List.json
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
