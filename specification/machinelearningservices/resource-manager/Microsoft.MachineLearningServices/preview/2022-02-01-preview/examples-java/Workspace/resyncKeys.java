import com.azure.core.util.Context;

/** Samples for Workspaces ResyncKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/resyncKeys.json
     */
    /**
     * Sample code: Resync Workspace Keys.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void resyncWorkspaceKeys(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaces().resyncKeys("testrg123", "workspaces123", Context.NONE);
    }
}
