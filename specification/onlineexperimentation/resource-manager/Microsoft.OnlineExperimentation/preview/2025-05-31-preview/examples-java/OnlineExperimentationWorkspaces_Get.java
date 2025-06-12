
/**
 * Samples for OnlineExperimentationWorkspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_Get.json
     */
    /**
     * Sample code: Get a single OnlineExperimentationWorkspace.
     * 
     * @param manager Entry point to OnlineExperimentationManager.
     */
    public static void getASingleOnlineExperimentationWorkspace(
        com.azure.resourcemanager.onlineexperimentation.OnlineExperimentationManager manager) {
        manager.onlineExperimentationWorkspaces().getByResourceGroupWithResponse("res9871", "expworkspace3",
            com.azure.core.util.Context.NONE);
    }
}
