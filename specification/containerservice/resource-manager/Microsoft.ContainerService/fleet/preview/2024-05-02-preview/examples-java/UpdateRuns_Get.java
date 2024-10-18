
/**
 * Samples for UpdateRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/
     * examples/UpdateRuns_Get.json
     */
    /**
     * Sample code: Gets an UpdateRun resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        getsAnUpdateRunResource(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().getWithResponse("rg1", "fleet1", "run1", com.azure.core.util.Context.NONE);
    }
}
