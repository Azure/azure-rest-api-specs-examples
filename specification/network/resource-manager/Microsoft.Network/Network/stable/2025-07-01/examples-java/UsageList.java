
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/UsageList.json
     */
    /**
     * Sample code: List usages.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listUsages(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getUsages().list("westus", com.azure.core.util.Context.NONE);
    }
}
