import com.azure.core.util.Context;

/** Samples for Compute ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Compute/listKeys.json
     */
    /**
     * Sample code: List AKS Compute Keys.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listAKSComputeKeys(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().listKeysWithResponse("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
