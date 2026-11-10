
/**
 * Samples for ScenarioTemplateVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ScenarioTemplateVersions_Get.json
     */
    /**
     * Sample code: Get a specific version of a built-in scenario template.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getASpecificVersionOfABuiltInScenarioTemplate(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioTemplateVersions().getWithResponse("ComputeZoneDown", "1.2", com.azure.core.util.Context.NONE);
    }
}
