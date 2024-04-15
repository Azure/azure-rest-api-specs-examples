
/**
 * Samples for FrontendEndpoints DisableHttps.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorDisableHttps.json
     */
    /**
     * Sample code: FrontendEndpoints_DisableHttps.
     * 
     * @param manager Entry point to FrontDoorManager.
     */
    public static void frontendEndpointsDisableHttps(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.frontendEndpoints().disableHttps("rg1", "frontDoor1", "frontendEndpoint1",
            com.azure.core.util.Context.NONE);
    }
}
