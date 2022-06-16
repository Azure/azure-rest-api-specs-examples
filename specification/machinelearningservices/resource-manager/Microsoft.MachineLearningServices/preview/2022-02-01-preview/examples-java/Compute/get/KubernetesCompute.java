import com.azure.core.util.Context;

/** Samples for Compute Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/get/KubernetesCompute.json
     */
    /**
     * Sample code: Get a Kubernetes Compute.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getAKubernetesCompute(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.computes().getWithResponse("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
