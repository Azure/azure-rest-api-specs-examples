/** Samples for UpdateRuns Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/UpdateRuns_Stop.json
     */
    /**
     * Sample code: Stops an UpdateRun.
     *
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void stopsAnUpdateRun(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().stop("rg1", "fleet1", "run1", null, com.azure.core.util.Context.NONE);
    }
}
