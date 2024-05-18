
/**
 * Samples for FleetMembers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2024-04-01/examples/
     * FleetMembers_Delete.json
     */
    /**
     * Sample code: Deletes a FleetMember resource asynchronously with a long running operation.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void deletesAFleetMemberResourceAsynchronouslyWithALongRunningOperation(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetMembers().delete("rg1", "fleet1", "member-1", null, com.azure.core.util.Context.NONE);
    }
}
