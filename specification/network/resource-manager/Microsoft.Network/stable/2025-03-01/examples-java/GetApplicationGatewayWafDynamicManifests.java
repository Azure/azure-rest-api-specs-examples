
/**
 * Samples for ApplicationGatewayWafDynamicManifests Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * GetApplicationGatewayWafDynamicManifests.json
     */
    /**
     * Sample code: Gets WAF manifests.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsWAFManifests(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGatewayWafDynamicManifests().get("westus",
            com.azure.core.util.Context.NONE);
    }
}
