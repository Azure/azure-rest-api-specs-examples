
/**
 * Samples for FleetUpdateStrategies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-02-02-preview/
     * examples/UpdateStrategies_Delete.json
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
