
/**
 * Samples for ExtensionMetadataV2 Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * extension/ExtensionMetadataV2_Get.json
     */
    /**
     * Sample code: GET an extension metadata.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAnExtensionMetadata(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.extensionMetadataV2s().getWithResponse("EastUS", "microsoft.azure.monitor", "azuremonitorlinuxagent",
            "1.33.0", com.azure.core.util.Context.NONE);
    }
}
