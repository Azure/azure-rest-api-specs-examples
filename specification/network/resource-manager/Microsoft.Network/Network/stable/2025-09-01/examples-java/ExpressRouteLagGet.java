
/**
 * Samples for ExpressRouteLags GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteLagGet.json
     */
    /**
     * Sample code: Get express route lag.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteLag(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags().getByResourceGroupWithResponse("rg1", "lagName",
            com.azure.core.util.Context.NONE);
    }
}
