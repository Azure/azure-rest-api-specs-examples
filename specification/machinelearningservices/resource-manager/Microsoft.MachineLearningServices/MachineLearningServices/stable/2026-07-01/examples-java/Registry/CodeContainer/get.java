
/**
 * Samples for RegistryCodeContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/CodeContainer/get.json
     */
    /**
     * Sample code: Get Registry Code Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryCodeContainers().getWithResponse("testrg123", "testregistry", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}
