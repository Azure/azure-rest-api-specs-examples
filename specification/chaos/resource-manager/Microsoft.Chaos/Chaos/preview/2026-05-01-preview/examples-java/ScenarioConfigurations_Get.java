
/**
 * Samples for ScenarioConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ScenarioConfigurations_Get.json
     */
    /**
     * Sample code: Get a scenario configuration.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAScenarioConfiguration(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioConfigurations().getWithResponse("exampleRG", "exampleWorkspace",
            "12345678-1234-1234-1234-123456789012", "config-5678-9012-3456-789012345678",
            com.azure.core.util.Context.NONE);
    }
}
