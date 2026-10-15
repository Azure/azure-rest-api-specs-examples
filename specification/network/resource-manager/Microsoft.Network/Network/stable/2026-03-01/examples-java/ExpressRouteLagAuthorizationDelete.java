
/**
 * Samples for ExpressRouteLagAuthorizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ExpressRouteLagAuthorizationDelete.json
     */
    /**
     * Sample code: Delete ExpressRouteLag Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteExpressRouteLagAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLagAuthorizations().delete("rg1", "expressRouteLagName",
            "authorizationName", com.azure.core.util.Context.NONE);
    }
}
