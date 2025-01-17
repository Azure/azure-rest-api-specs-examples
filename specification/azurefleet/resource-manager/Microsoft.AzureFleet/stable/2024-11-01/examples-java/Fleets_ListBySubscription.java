
/**
 * Samples for Fleets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Fleets_ListBySubscription.json
     */
    /**
     * Sample code: Fleets_ListBySubscription.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsListBySubscription(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().list(com.azure.core.util.Context.NONE);
    }
}
