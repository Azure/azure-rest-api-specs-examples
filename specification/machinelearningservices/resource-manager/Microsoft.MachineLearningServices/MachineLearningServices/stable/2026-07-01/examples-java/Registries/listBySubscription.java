
/**
 * Samples for Registries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registries/listBySubscription.json
     */
    /**
     * Sample code: List registries by subscription.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistriesBySubscription(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registries().list(com.azure.core.util.Context.NONE);
    }
}
