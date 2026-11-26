
/**
 * Samples for RegistryModelContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/ModelContainer/get.json
     */
    /**
     * Sample code: Get Registry Model Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryModelContainers().getWithResponse("testrg123", "registry123", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}
