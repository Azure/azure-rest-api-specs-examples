
/**
 * Samples for Fleets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2024-04-01/examples/
     * Fleets_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists the Fleet resources in a resource group.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listsTheFleetResourcesInAResourceGroup(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleets().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
