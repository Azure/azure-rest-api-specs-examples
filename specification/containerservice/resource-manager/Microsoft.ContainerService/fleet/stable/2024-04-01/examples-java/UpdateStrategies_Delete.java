
/**
 * Samples for FleetUpdateStrategies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2024-04-01/examples/
     * UpdateStrategies_Delete.json
     */
    /**
     * Sample code: Delete a FleetUpdateStrategy resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void deleteAFleetUpdateStrategyResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetUpdateStrategies().delete("rg1", "fleet1", "strategy1", null, com.azure.core.util.Context.NONE);
    }
}
