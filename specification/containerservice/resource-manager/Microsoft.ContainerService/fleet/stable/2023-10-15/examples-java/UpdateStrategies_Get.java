/** Samples for FleetUpdateStrategies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/UpdateStrategies_Get.json
     */
    /**
     * Sample code: Get a FleetUpdateStrategy resource.
     *
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void getAFleetUpdateStrategyResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetUpdateStrategies().getWithResponse("rg1", "fleet1", "strategy1", com.azure.core.util.Context.NONE);
    }
}
