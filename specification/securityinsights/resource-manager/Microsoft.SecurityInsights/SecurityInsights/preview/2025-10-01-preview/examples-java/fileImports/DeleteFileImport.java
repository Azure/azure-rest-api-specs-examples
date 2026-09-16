
/**
 * Samples for FileImports Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/fileImports/DeleteFileImport.json
     */
    /**
     * Sample code: Delete a file import.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAFileImport(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.fileImports().delete("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}
