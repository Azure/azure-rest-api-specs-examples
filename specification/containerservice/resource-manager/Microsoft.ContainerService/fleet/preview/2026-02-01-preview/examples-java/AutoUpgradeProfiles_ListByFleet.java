
/**
 * Samples for AutoUpgradeProfiles ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/AutoUpgradeProfiles_ListByFleet.json
     */
    /**
     * Sample code: Lists the AutoUpgradeProfile resources by fleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listsTheAutoUpgradeProfileResourcesByFleet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.autoUpgradeProfiles().listByFleet("rgfleets", "fleet1", null, null, com.azure.core.util.Context.NONE);
    }
}
