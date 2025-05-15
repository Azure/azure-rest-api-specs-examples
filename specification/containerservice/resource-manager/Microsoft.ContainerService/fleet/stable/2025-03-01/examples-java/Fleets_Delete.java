
/**
 * Samples for Fleets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Fleets_Delete.json
     */
    /**
     * Sample code: Deletes a Fleet resource asynchronously with a long running operation.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void deletesAFleetResourceAsynchronouslyWithALongRunningOperation(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleets().delete("rg1", "fleet1", null, com.azure.core.util.Context.NONE);
    }
}
