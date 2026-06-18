
/**
 * Samples for ScenarioConfigurations FixResourcePermissions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ScenarioConfigurations_FixResourcePermissions.json
     */
    /**
     * Sample code: Fixes resource permissions for the given scenario configuration.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        fixesResourcePermissionsForTheGivenScenarioConfiguration(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioConfigurations().fixResourcePermissions("exampleRG", "exampleWorkspace",
            "12345678-1234-1234-1234-123456789012", "config-5678-9012-3456-789012345678", null,
            com.azure.core.util.Context.NONE);
    }
}
