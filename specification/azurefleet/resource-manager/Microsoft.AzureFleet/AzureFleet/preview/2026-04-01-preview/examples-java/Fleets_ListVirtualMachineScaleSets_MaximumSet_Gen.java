
/**
 * Samples for Fleets ListVirtualMachineScaleSets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/Fleets_ListVirtualMachineScaleSets_MaximumSet_Gen.json
     */
    /**
     * Sample code: Fleets_ListVirtualMachineScaleSets_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsListVirtualMachineScaleSetsMaximumSetGen(
        com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().listVirtualMachineScaleSets("rgazurefleet", "myFleet", com.azure.core.util.Context.NONE);
    }
}
