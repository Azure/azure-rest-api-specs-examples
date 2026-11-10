
/**
 * Samples for ScenarioTemplates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ScenarioTemplates_Get.json
     */
    /**
     * Sample code: Get the latest version of a built-in scenario template.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getTheLatestVersionOfABuiltInScenarioTemplate(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioTemplates().getWithResponse("ComputeZoneDown", com.azure.core.util.Context.NONE);
    }
}
