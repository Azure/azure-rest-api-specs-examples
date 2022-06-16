import com.azure.core.util.Context;

/** Samples for Workspaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/listBySubscription.json
     */
    /**
     * Sample code: Get Workspaces by subscription.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getWorkspacesBySubscription(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaces().list(null, Context.NONE);
    }
}
