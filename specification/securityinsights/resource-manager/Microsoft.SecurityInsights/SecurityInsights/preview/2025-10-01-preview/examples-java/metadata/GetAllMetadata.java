
/**
 * Samples for Metadata List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/metadata/GetAllMetadata.json
     */
    /**
     * Sample code: Get all metadata.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllMetadata(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.metadatas().list("myRg", "myWorkspace", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
