
/**
 * Samples for Fleets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/Fleets_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Fleets_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsGetMaximumSetGen(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().getByResourceGroupWithResponse("rgazurefleet", "myFleet", com.azure.core.util.Context.NONE);
    }
}
