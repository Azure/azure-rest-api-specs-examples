
/**
 * Samples for AutoUpgradeProfileOperations GenerateUpdateRun.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/AutoUpgradeProfileOperations_GenerateUpdateRun.json
     */
    /**
     * Sample code: AutoUpgradeProfileOperations_GenerateUpdateRun.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void autoUpgradeProfileOperationsGenerateUpdateRun(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.autoUpgradeProfileOperations().generateUpdateRun("rgfleets", "fleet1", "aup1",
            com.azure.core.util.Context.NONE);
    }
}
