
/**
 * Samples for UpdateRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/UpdateRuns_Get.json
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
