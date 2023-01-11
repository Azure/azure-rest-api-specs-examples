import com.azure.core.util.Context;

/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Usage/list.json
     */
    /**
     * Sample code: List Usages.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listUsages(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.usages().list("eastus", Context.NONE);
    }
}
