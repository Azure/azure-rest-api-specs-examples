/** Samples for FrontendEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorFrontendEndpointGet.json
     */
    /**
     * Sample code: Get Frontend Endpoint.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getFrontendEndpoint(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager
            .frontendEndpoints()
            .getWithResponse("rg1", "frontDoor1", "frontendEndpoint1", com.azure.core.util.Context.NONE);
    }
}
