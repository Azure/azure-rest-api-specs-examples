
/**
 * Samples for Registries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registries/delete.json
     */
    /**
     * Sample code: Delete Registry.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteRegistry(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registries().delete("test-rg", "string", com.azure.core.util.Context.NONE);
    }
}
