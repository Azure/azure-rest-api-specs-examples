import com.azure.core.util.Context;

/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Usage/list.json
     */
    /**
     * Sample code: List Usages.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listUsages(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.usages().list("eastus", Context.NONE);
    }
}
