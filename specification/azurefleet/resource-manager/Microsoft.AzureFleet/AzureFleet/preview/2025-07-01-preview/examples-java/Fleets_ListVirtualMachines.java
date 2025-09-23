
/**
 * Samples for Fleets ListVirtualMachines.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/Fleets_ListVirtualMachines.json
     */
    /**
     * Sample code: Fleets_ListVirtualMachines_MaximumSet.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void
        fleetsListVirtualMachinesMaximumSet(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().listVirtualMachines("rgazurefleet", "myFleet", "xzcepyottghqa", "hydepbmwuypaprlphcdecsz",
            com.azure.core.util.Context.NONE);
    }
}
