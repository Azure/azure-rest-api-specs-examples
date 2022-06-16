import com.azure.core.util.Context;

/** Samples for Compute ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/listKeys.json
     */
    /**
     * Sample code: List AKS Compute Keys.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listAKSComputeKeys(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.computes().listKeysWithResponse("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
