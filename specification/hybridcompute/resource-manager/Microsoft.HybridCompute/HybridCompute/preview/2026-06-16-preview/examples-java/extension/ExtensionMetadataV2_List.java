
/**
 * Samples for ExtensionMetadataV2 List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/extension/ExtensionMetadataV2_List.json
     */
    /**
     * Sample code: GET a list of extension metadata.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        gETAListOfExtensionMetadata(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.extensionMetadataV2s().list("EastUS", "microsoft.azure.monitor", "azuremonitorlinuxagent",
            com.azure.core.util.Context.NONE);
    }
}
