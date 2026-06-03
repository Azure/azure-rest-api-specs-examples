
/**
 * Samples for FleetMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/FleetMembers_Get.json
     */
    /**
     * Sample code: Gets a FleetMember resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        getsAFleetMemberResource(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetMembers().getWithResponse("rgfleets", "fleet1", "member1", com.azure.core.util.Context.NONE);
    }
}
