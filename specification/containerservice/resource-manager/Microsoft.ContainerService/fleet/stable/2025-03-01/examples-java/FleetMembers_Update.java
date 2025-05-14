
import com.azure.resourcemanager.containerservicefleet.models.FleetMember;

/**
 * Samples for FleetMembers UpdateAsync.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/FleetMembers_Update.json
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
