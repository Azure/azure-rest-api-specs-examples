
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/ListUsagesWithResourceMetadata.json
     */
    /**
     * Sample code: Get Usages With Resource Metadata.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getUsagesWithResourceMetadata(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.usages().list("WestUS", null, com.azure.core.util.Context.NONE);
    }
}
