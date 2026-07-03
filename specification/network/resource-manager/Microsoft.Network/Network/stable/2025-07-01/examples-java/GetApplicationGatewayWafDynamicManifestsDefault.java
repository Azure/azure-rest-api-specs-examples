
/**
 * Samples for ApplicationGatewayWafDynamicManifestsDefault Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/GetApplicationGatewayWafDynamicManifestsDefault.json
     */
    /**
     * Sample code: Gets WAF default manifest.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsWAFDefaultManifest(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGatewayWafDynamicManifestsDefaults().getWithResponse("westus",
            com.azure.core.util.Context.NONE);
    }
}
