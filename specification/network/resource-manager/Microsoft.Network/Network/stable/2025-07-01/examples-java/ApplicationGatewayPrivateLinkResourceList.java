
/**
 * Samples for ApplicationGatewayPrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayPrivateLinkResourceList.json
     */
    /**
     * Sample code: Lists all private link resources on application gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listsAllPrivateLinkResourcesOnApplicationGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGatewayPrivateLinkResources().list("rg1", "appgw",
            com.azure.core.util.Context.NONE);
    }
}
