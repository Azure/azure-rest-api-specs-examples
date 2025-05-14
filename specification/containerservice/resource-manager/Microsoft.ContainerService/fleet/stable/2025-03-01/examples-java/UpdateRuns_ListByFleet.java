
/**
 * Samples for UpdateRuns ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/UpdateRuns_ListByFleet.json
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
