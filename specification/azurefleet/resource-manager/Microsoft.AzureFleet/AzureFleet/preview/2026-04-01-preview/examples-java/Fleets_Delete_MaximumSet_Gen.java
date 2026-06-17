
/**
 * Samples for Fleets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/Fleets_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Fleets_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsDeleteMaximumSetGen(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().delete("rgazurefleet", "testFleet", com.azure.core.util.Context.NONE);
    }
}
