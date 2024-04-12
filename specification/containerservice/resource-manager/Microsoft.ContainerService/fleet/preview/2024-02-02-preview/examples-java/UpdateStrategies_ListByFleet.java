
/**
 * Samples for FleetUpdateStrategies ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-02-02-preview/
     * examples/UpdateStrategies_ListByFleet.json
     */
    /**
     * Sample code: List the FleetUpdateStrategy resources by fleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listTheFleetUpdateStrategyResourcesByFleet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetUpdateStrategies().listByFleet("rg1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
