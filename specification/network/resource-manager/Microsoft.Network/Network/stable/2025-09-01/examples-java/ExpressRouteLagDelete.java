
/**
 * Samples for ExpressRouteLags Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteLagDelete.json
     */
    /**
     * Sample code: Delete express route lag.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteExpressRouteLag(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags().delete("rg1", "lagName", com.azure.core.util.Context.NONE);
    }
}
