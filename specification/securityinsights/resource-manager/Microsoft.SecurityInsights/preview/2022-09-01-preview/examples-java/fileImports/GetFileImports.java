import com.azure.core.util.Context;

/** Samples for FileImports List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/fileImports/GetFileImports.json
     */
    /**
     * Sample code: Get all file imports.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllFileImports(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .fileImports()
            .list("myRg", "myWorkspace", null, "properties/createdTimeUtc desc", 1, null, Context.NONE);
    }
}
