
/**
 * Samples for ExtensionMetadata Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/
     * extension/ExtensionMetadata_Get.json
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
