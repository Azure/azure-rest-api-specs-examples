
/**
 * Samples for ApplicationGatewayWafDynamicManifests Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/GetApplicationGatewayWafDynamicManifests.json
     */
    /**
     * Sample code: Gets WAF manifests.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsWAFManifests(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGatewayWafDynamicManifests().get("westus",
            com.azure.core.util.Context.NONE);
    }
}
