
/**
 * Samples for Metadata Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/metadata/GetMetadata.json
     */
    /**
     * Sample code: Get single metadata by name.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getSingleMetadataByName(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.metadatas().getWithResponse("myRg", "myWorkspace", "metadataName", com.azure.core.util.Context.NONE);
    }
}
