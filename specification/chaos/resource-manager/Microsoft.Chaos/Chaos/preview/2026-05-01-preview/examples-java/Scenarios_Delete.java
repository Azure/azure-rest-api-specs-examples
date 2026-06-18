
/**
 * Samples for Scenarios Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Scenarios_Delete.json
     */
    /**
     * Sample code: Delete a Scenario in a workspace.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void deleteAScenarioInAWorkspace(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarios().deleteWithResponse("exampleRG", "exampleWorkspace", "myVMShutdownScenario",
            com.azure.core.util.Context.NONE);
    }
}
