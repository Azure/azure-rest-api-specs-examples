
/**
 * Samples for Fleets ListCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2024-04-01/examples/
     * Fleets_ListCredentialsResult.json
     */
    /**
     * Sample code: Lists the user credentials of a Fleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listsTheUserCredentialsOfAFleet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleets().listCredentialsWithResponse("rg1", "fleet", com.azure.core.util.Context.NONE);
    }
}
