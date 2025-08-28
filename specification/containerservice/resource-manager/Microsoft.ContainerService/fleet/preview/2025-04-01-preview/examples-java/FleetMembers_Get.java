
/**
 * Samples for FleetMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/FleetMembers_Get.json
     */
    /**
     * Sample code: Gets a FleetMember resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        getsAFleetMemberResource(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetMembers().getWithResponse("rg1", "fleet1", "member-1", com.azure.core.util.Context.NONE);
    }
}
