
/**
 * Samples for Fleets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Fleets_ListBySub.json
     */
    /**
     * Sample code: Lists the Fleet resources in a subscription.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listsTheFleetResourcesInASubscription(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleets().list(com.azure.core.util.Context.NONE);
    }
}
