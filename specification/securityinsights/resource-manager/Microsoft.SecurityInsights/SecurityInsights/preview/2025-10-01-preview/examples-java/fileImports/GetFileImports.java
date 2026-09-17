
/**
 * Samples for FileImports List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/fileImports/GetFileImports.json
     */
    /**
     * Sample code: Get all file imports.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllFileImports(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.fileImports().list("myRg", "myWorkspace", null, "properties/createdTimeUtc desc", 1, null,
            com.azure.core.util.Context.NONE);
    }
}
