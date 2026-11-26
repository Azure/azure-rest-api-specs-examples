
/**
 * Samples for OnlineDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/OnlineDeployment/delete.json
     */
    /**
     * Sample code: Delete Workspace Online Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceOnlineDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineDeployments().delete("testrg123", "workspace123", "testEndpoint", "testDeployment",
            com.azure.core.util.Context.NONE);
    }
}
