
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Usage/list.json
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
