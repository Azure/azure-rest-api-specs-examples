
/**
 * Samples for Endpoints ListIngressGatewayCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/EndpointsPostListIngressGatewayCredentials.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsPostListIngressGatewayCredentials.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsPostListIngressGatewayCredentials(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.endpoints().listIngressGatewayCredentialsWithResponse(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/arcGroup/providers/Microsoft.ArcPlaceHolder/ProvisionedClusters/cluster0",
            "default", 10800L, null, com.azure.core.util.Context.NONE);
    }
}
