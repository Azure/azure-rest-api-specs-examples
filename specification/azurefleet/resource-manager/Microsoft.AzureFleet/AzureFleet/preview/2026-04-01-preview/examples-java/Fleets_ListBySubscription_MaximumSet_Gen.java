
/**
 * Samples for Fleets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/Fleets_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Fleets_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void
        fleetsListBySubscriptionMaximumSetGen(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().list(com.azure.core.util.Context.NONE);
    }
}
