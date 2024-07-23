
/**
 * Samples for Fleets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azurefleet/AzureFleet.Management/examples/2024-05-01-preview/Fleets_Delete.json
     */
    /**
     * Sample code: Fleets_Delete.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsDelete(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().delete("rgazurefleet", "testFleet", com.azure.core.util.Context.NONE);
    }
}
