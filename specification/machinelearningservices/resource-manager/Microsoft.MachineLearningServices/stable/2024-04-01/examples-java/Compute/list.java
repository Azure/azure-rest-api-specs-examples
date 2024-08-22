
/**
 * Samples for Compute List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/list.json
     */
    /**
     * Sample code: Get Computes.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getComputes(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().list("testrg123", "workspaces123", null, com.azure.core.util.Context.NONE);
    }
}
