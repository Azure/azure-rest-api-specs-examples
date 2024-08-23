
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Usage/list.json
     */
    /**
     * Sample code: List Usages.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listUsages(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.usages().list("eastus", com.azure.core.util.Context.NONE);
    }
}
