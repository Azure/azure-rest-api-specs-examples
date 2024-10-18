
/**
 * Samples for AutoUpgradeProfiles ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/
     * examples/AutoUpgradeProfiles_ListByFleet.json
     */
    /**
     * Sample code: Lists the AutoUpgradeProfile resources by fleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listsTheAutoUpgradeProfileResourcesByFleet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.autoUpgradeProfiles().listByFleet("rg1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
