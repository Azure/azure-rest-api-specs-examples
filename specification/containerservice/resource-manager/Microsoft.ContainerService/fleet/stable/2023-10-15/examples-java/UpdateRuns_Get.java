/** Samples for UpdateRuns Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/UpdateRuns_Get.json
     */
    /**
     * Sample code: Gets an UpdateRun resource.
     *
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void getsAnUpdateRunResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().getWithResponse("rg1", "fleet1", "run1", com.azure.core.util.Context.NONE);
    }
}
