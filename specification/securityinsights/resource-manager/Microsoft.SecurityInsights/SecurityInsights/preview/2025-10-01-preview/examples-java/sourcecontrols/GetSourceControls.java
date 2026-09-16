
/**
 * Samples for SourceControls List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/sourcecontrols/GetSourceControls.json
     */
    /**
     * Sample code: Get all source controls.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllSourceControls(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sourceControls().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
