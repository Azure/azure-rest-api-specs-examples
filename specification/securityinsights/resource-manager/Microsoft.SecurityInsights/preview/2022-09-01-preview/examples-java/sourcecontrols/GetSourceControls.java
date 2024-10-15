
/**
 * Samples for SourceControlsOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * sourcecontrols/GetSourceControls.json
     */
    /**
     * Sample code: Get all source controls.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllSourceControls(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sourceControlsOperations().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
