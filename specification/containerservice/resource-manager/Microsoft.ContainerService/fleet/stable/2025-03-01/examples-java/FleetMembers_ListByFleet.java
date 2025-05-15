
/**
 * Samples for FleetMembers ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/FleetMembers_ListByFleet.json
     */
    /**
     * Sample code: Lists the members of a Fleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        listsTheMembersOfAFleet(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetMembers().listByFleet("rg1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
