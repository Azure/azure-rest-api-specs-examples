
/**
 * Samples for ExtensionMetadata List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/
     * extension/ExtensionMetadata_List.json
     */
    /**
     * Sample code: GET a list of extensions.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAListOfExtensions(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.extensionMetadatas().list("EastUS", "microsoft.azure.monitor", "azuremonitorlinuxagent",
            com.azure.core.util.Context.NONE);
    }
}
