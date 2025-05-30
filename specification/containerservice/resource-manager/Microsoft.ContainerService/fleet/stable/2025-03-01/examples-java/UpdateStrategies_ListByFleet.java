
/**
 * Samples for FleetUpdateStrategies ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/UpdateStrategies_ListByFleet.json
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
