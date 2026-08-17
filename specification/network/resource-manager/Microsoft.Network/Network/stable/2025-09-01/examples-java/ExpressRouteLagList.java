
/**
 * Samples for ExpressRouteLags List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteLagList.json
     */
    /**
     * Sample code: List all express route lags.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllExpressRouteLags(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags().list(com.azure.core.util.Context.NONE);
    }
}
