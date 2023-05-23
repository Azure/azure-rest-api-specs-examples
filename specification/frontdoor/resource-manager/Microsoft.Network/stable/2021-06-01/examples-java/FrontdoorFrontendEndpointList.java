/** Samples for FrontendEndpoints ListByFrontDoor. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorFrontendEndpointList.json
     */
    /**
     * Sample code: List Frontend endpoints in a Front Door.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void listFrontendEndpointsInAFrontDoor(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.frontendEndpoints().listByFrontDoor("rg1", "frontDoor1", com.azure.core.util.Context.NONE);
    }
}
