import com.azure.core.util.Context;

/** Samples for Workspaces ListNotebookAccessToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/listNotebookAccessToken.json
     */
    /**
     * Sample code: List Workspace Keys.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listWorkspaceKeys(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaces().listNotebookAccessTokenWithResponse("workspace-1234", "testworkspace", Context.NONE);
    }
}
