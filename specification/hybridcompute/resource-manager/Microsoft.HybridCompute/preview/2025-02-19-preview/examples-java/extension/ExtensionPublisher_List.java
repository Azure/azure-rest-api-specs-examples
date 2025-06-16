
/**
 * Samples for ExtensionPublisher List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * extension/ExtensionPublisher_List.json
     */
    /**
     * Sample code: GET a list of extension publishers.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        gETAListOfExtensionPublishers(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.extensionPublishers().list("EastUS", com.azure.core.util.Context.NONE);
    }
}
