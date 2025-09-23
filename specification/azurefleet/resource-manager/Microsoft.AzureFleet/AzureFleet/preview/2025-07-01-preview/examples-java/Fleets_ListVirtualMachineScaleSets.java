
/**
 * Samples for Fleets ListVirtualMachineScaleSets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/Fleets_ListVirtualMachineScaleSets.json
     */
    /**
     * Sample code: Fleets_ListVirtualMachineScaleSets.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void
        fleetsListVirtualMachineScaleSets(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().listVirtualMachineScaleSets("rgazurefleet", "myFleet", com.azure.core.util.Context.NONE);
    }
}
