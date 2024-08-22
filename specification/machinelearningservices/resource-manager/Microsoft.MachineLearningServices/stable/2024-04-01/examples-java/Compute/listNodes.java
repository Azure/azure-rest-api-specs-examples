
/**
 * Samples for Compute ListNodes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/listNodes.json
     */
    /**
     * Sample code: Get compute nodes information for a compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getComputeNodesInformationForACompute(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().listNodes("testrg123", "workspaces123", "compute123", com.azure.core.util.Context.NONE);
    }
}
