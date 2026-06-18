
/**
 * Samples for Scenarios Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Scenarios_Get.json
     */
    /**
     * Sample code: Get a scenario.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAScenario(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarios().getWithResponse("exampleRG", "exampleWorkspace", "zoneDownScenario",
            com.azure.core.util.Context.NONE);
    }
}
