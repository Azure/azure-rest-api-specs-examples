
/**
 * Samples for ExpressRouteLags LinksGet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteLagLinkGet.json
     */
    /**
     * Sample code: Get express route lag link.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteLagLink(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags().linksGetWithResponse("rg1", "lagName", "linkName",
            com.azure.core.util.Context.NONE);
    }
}
