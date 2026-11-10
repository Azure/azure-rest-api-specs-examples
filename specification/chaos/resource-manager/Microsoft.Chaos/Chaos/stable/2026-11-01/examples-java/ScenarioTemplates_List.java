
/**
 * Samples for ScenarioTemplates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ScenarioTemplates_List.json
     */
    /**
     * Sample code: List the built-in scenario templates.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listTheBuiltInScenarioTemplates(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioTemplates().list(com.azure.core.util.Context.NONE);
    }
}
