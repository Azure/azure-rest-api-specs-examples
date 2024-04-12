
/**
 * Samples for FleetMembers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-02-02-preview/
     * examples/FleetMembers_Delete.json
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
