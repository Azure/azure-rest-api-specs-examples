
/**
 * Samples for ExpressRouteLags LinksList.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteLagLinkList.json
     */
    /**
     * Sample code: List express route lag links.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteLagLinks(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags().linksList("rg1", "lagName", com.azure.core.util.Context.NONE);
    }
}
