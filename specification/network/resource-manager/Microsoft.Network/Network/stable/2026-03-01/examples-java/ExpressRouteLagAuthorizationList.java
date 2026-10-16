
/**
 * Samples for ExpressRouteLagAuthorizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ExpressRouteLagAuthorizationList.json
     */
    /**
     * Sample code: List ExpressRouteLag Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteLagAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLagAuthorizations().list("rg1", "expressRouteLagName",
            com.azure.core.util.Context.NONE);
    }
}
