
/**
 * Samples for Fleets ListVirtualMachines.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/Fleets_ListVirtualMachines_MaximumSet_Gen.json
     */
    /**
     * Sample code: Fleets_ListVirtualMachines_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void
        fleetsListVirtualMachinesMaximumSetGen(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().listVirtualMachines("rgazurefleet", "testFleet", "qppsnaauhedxu", "jxgpugummyphgx",
            com.azure.core.util.Context.NONE);
    }
}
