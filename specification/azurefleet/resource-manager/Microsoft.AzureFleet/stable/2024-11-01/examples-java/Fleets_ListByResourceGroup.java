
/**
 * Samples for Fleets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Fleets_ListByResourceGroup.json
     */
    /**
     * Sample code: Fleets_ListByResourceGroup.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsListByResourceGroup(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().listByResourceGroup("rgazurefleet", com.azure.core.util.Context.NONE);
    }
}
