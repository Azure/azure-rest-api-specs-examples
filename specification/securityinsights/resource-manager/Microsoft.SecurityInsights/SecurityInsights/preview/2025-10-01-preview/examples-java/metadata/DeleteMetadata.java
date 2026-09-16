
/**
 * Samples for Metadata Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/metadata/DeleteMetadata.json
     */
    /**
     * Sample code: Delete metadata.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteMetadata(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.metadatas().deleteWithResponse("myRg", "myWorkspace", "metadataName", com.azure.core.util.Context.NONE);
    }
}
