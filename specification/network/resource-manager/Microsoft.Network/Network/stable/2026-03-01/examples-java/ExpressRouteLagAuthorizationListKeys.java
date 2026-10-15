
/**
 * Samples for ExpressRouteLagAuthorizations ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ExpressRouteLagAuthorizationListKeys.json
     */
    /**
     * Sample code: List ExpressRouteLag Authorization Keys.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteLagAuthorizationKeys(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLagAuthorizations().listKeysWithResponse("rg1", "expressRouteLagName",
            "authorizationName", com.azure.core.util.Context.NONE);
    }
}
