
/**
 * Samples for ExtensionMetadata Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/extension/ExtensionMetadata_Get.json
     */
    /**
     * Sample code: GET an extensions metadata.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAnExtensionsMetadata(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.extensionMetadatas().getWithResponse("EastUS", "microsoft.azure.monitor", "azuremonitorlinuxagent",
            "1.9.1", com.azure.core.util.Context.NONE);
    }
}
