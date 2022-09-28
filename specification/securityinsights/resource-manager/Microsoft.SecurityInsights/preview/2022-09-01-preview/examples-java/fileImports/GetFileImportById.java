import com.azure.core.util.Context;

/** Samples for FileImports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/fileImports/GetFileImportById.json
     */
    /**
     * Sample code: Get a file import.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAFileImport(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .fileImports()
            .getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", Context.NONE);
    }
}
