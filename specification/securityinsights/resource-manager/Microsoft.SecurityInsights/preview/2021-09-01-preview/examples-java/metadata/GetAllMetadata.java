import com.azure.core.util.Context;

/** Samples for Metadata List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/metadata/GetAllMetadata.json
     */
    /**
     * Sample code: Get all metadata.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllMetadata(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.metadatas().list("myRg", "myWorkspace", null, null, null, null, Context.NONE);
    }
}
