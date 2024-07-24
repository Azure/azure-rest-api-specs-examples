
/**
 * Samples for Fleets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurefleet/AzureFleet.Management/examples/2024-05-01-preview/Fleets_ListByResourceGroup.json
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
