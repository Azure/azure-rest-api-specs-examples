
/**
 * Samples for SourceControls Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/sourcecontrols/GetSourceControlById.json
     */
    /**
     * Sample code: Get a source control.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getASourceControl(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sourceControls().getWithResponse("myRg", "myWorkspace", "789e0c1f-4a3d-43ad-809c-e713b677b04a",
            com.azure.core.util.Context.NONE);
    }
}
