
/**
 * Samples for SourceControlsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * sourcecontrols/GetSourceControlById.json
     */
    /**
     * Sample code: Get a source control.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getASourceControl(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sourceControlsOperations().getWithResponse("myRg", "myWorkspace",
            "789e0c1f-4a3d-43ad-809c-e713b677b04a", com.azure.core.util.Context.NONE);
    }
}
