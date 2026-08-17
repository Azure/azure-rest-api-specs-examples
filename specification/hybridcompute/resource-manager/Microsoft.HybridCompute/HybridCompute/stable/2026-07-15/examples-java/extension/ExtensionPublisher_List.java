
/**
 * Samples for ExtensionPublisher List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/extension/ExtensionPublisher_List.json
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
