import com.azure.core.util.Context;

/** Samples for WorkspaceFeatures List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceFeature/list.json
     */
    /**
     * Sample code: List Workspace features.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listWorkspaceFeatures(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaceFeatures().list("myResourceGroup", "testworkspace", Context.NONE);
    }
}
