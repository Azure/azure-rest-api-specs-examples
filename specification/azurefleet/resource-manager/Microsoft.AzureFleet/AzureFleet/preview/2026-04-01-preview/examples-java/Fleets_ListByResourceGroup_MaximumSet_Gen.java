
/**
 * Samples for Fleets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/Fleets_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Fleets_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void
        fleetsListByResourceGroupMaximumSetGen(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().listByResourceGroup("rgazurefleet", com.azure.core.util.Context.NONE);
    }
}
