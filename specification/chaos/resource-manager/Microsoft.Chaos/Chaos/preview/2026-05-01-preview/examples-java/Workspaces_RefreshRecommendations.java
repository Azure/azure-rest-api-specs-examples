
/**
 * Samples for Workspaces RefreshRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Workspaces_RefreshRecommendations.json
     */
    /**
     * Sample code: Refresh recommendations for all scenarios in a workspace.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        refreshRecommendationsForAllScenariosInAWorkspace(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().refreshRecommendations("exampleRG", "exampleWorkspace", com.azure.core.util.Context.NONE);
    }
}
