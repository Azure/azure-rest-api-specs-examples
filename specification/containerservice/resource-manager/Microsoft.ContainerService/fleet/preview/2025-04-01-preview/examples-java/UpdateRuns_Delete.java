
/**
 * Samples for UpdateRuns Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/UpdateRuns_Delete.json
     */
    /**
     * Sample code: Delete an updateRun resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void deleteAnUpdateRunResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().delete("rg1", "fleet1", "run1", null, com.azure.core.util.Context.NONE);
    }
}
