
/**
 * Samples for Workspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/listBySubscription.json
     */
    /**
     * Sample code: Get Workspaces by subscription.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspacesBySubscription(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().list(null, com.azure.core.util.Context.NONE);
    }
}
