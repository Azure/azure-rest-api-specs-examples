import com.azure.core.util.Context;

/** Samples for Workspaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/get.json
     */
    /**
     * Sample code: Get Workspace.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getWorkspace(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("workspace-1234", "testworkspace", Context.NONE);
    }
}
