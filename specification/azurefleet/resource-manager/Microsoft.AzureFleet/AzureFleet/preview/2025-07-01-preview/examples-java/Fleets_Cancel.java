
/**
 * Samples for Fleets Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/Fleets_Cancel.json
     */
    /**
     * Sample code: Fleets_Cancel.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsCancel(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().cancel("rgazurefleet", "myFleet", com.azure.core.util.Context.NONE);
    }
}
