/** Samples for FleetMembers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-06-15-preview/examples/FleetMembers_Get.json
     */
    /**
     * Sample code: Gets a FleetMember resource.
     *
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void getsAFleetMemberResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetMembers().getWithResponse("rg1", "fleet1", "member-1", com.azure.core.util.Context.NONE);
    }
}
