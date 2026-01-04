
/**
 * Samples for ApplicationGatewayWafDynamicManifestsDefault Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * GetApplicationGatewayWafDynamicManifestsDefault.json
     */
    /**
     * Sample code: Gets WAF default manifest.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsWAFDefaultManifest(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGatewayWafDynamicManifestsDefaults()
            .getWithResponse("westus", com.azure.core.util.Context.NONE);
    }
}
