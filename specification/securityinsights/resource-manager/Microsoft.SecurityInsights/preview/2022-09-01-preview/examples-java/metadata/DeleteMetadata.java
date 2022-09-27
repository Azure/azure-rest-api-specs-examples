import com.azure.core.util.Context;

/** Samples for Metadata Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/DeleteMetadata.json
     */
    /**
     * Sample code: Delete metadata.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteMetadata(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.metadatas().deleteWithResponse("myRg", "myWorkspace", "metadataName", Context.NONE);
    }
}
