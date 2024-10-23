
/**
 * Samples for Fleets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Fleets_Delete.json
     */
    /**
     * Sample code: Fleets_Delete.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsDelete(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().delete("rgazurefleet", "testFleet", com.azure.core.util.Context.NONE);
    }
}
