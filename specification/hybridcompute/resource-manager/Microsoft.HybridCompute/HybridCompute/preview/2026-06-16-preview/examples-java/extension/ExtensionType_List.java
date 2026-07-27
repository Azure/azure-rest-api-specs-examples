
/**
 * Samples for ExtensionType List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/extension/ExtensionType_List.json
     */
    /**
     * Sample code: GET a list of extension types.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAListOfExtensionTypes(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.extensionTypes().list("EastUS", "microsoft.azure.monitor", com.azure.core.util.Context.NONE);
    }
}
