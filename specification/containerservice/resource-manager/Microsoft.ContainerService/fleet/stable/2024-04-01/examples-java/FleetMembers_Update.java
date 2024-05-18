
import com.azure.resourcemanager.containerservicefleet.models.FleetMember;

/**
 * Samples for FleetMembers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2024-04-01/examples/
     * FleetMembers_Update.json
     */
    /**
     * Sample code: Updates a FleetMember resource synchronously.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void updatesAFleetMemberResourceSynchronously(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        FleetMember resource = manager.fleetMembers()
            .getWithResponse("rg1", "fleet1", "member-1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withGroup("staging").apply();
    }
}
