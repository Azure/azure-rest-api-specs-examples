
/**
 * Samples for FleetMembers Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/
     * examples/FleetMembers_Create.json
     */
    /**
     * Sample code: Creates a FleetMember resource with a long running operation.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void createsAFleetMemberResourceWithALongRunningOperation(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetMembers().define("member-1").withExistingFleet("rg1", "fleet1").withClusterResourceId(
            "/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster-1")
            .create();
    }
}
