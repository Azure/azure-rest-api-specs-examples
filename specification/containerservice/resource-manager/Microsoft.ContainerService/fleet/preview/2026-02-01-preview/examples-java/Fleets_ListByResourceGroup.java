
/**
 * Samples for Fleets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/Fleets_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists the Fleet resources in a resource group.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listsTheFleetResourcesInAResourceGroup(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleets().listByResourceGroup("rgfleets", com.azure.core.util.Context.NONE);
    }
}
