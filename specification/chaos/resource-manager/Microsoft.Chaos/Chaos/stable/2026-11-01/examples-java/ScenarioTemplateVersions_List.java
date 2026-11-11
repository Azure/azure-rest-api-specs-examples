
/**
 * Samples for ScenarioTemplateVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ScenarioTemplateVersions_List.json
     */
    /**
     * Sample code: List the published versions of a built-in scenario template.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        listThePublishedVersionsOfABuiltInScenarioTemplate(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioTemplateVersions().list("ComputeZoneDown", com.azure.core.util.Context.NONE);
    }
}
