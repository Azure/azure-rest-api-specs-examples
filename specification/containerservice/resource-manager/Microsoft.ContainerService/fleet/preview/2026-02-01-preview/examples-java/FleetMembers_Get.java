
/**
 * Samples for FleetMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/FleetMembers_Get.json
     */
    /**
     * Sample code: Gets a FleetMember resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        getsAFleetMemberResource(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetMembers().getWithResponse("rgfleets", "fleet1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
