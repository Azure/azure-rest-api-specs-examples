
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/Operations_List.
     * json
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
