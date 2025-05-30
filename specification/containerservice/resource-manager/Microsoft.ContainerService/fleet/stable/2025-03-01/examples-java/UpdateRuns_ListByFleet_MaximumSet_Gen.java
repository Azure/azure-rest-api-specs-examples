
/**
 * Samples for UpdateRuns ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/UpdateRuns_ListByFleet_MaximumSet_Gen.json
     */
    /**
     * Sample code: Lists the UpdateRun resources by fleet. - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listsTheUpdateRunResourcesByFleetGeneratedByMaximumSetRule(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().listByFleet("rgfleets", "fleet1", com.azure.core.util.Context.NONE);
    }
}
