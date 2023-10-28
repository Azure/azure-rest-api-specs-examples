/** Samples for UpdateRuns ListByFleet. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/UpdateRuns_ListByFleet.json
     */
    /**
     * Sample code: Lists the UpdateRun resources by fleet.
     *
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listsTheUpdateRunResourcesByFleet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().listByFleet("rg1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
