
/**
 * Samples for Metadata List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/metadata/GetAllMetadataOData.json
     */
    /**
     * Sample code: Get all metadata with OData filter/orderby/skip/top.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllMetadataWithODataFilterOrderbySkipTop(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.metadatas().list("myRg", "myWorkspace", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
