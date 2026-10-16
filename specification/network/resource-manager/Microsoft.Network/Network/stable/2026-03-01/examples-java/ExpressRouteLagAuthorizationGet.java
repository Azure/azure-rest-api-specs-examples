
/**
 * Samples for ExpressRouteLagAuthorizations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ExpressRouteLagAuthorizationGet.json
     */
    /**
     * Sample code: Get ExpressRouteLag Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteLagAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLagAuthorizations().getWithResponse("rg1", "expressRouteLagName",
            "authorizationName", com.azure.core.util.Context.NONE);
    }
}
