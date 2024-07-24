
/**
 * Samples for Fleets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azurefleet/AzureFleet.Management/examples/2024-05-01-preview/Fleets_Get.json
     */
    /**
     * Sample code: Fleets_Get.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void fleetsGet(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().getByResourceGroupWithResponse("rgazurefleet", "testFleet", com.azure.core.util.Context.NONE);
    }
}
