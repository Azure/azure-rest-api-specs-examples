/** Samples for UpdateRuns Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/UpdateRuns_Start.json
     */
    /**
     * Sample code: Starts an UpdateRun.
     *
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void startsAnUpdateRun(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().start("rg1", "fleet1", "run1", null, com.azure.core.util.Context.NONE);
    }
}
