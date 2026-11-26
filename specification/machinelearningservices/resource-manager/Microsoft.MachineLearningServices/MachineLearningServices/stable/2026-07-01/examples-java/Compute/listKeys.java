
/**
 * Samples for Compute ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Compute/listKeys.json
     */
    /**
     * Sample code: List AKS Compute Keys.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listAKSComputeKeys(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().listKeysWithResponse("testrg123", "workspaces123", "compute123",
            com.azure.core.util.Context.NONE);
    }
}
