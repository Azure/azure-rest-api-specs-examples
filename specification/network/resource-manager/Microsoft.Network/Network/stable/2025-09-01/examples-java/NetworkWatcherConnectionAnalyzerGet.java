
/**
 * Samples for NetworkWatchers ConnectionAnalyzersGet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkWatcherConnectionAnalyzerGet.json
     */
    /**
     * Sample code: Get connection analyzer.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getConnectionAnalyzer(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().connectionAnalyzersGetWithResponse("connectionAnalyzerRG", "nw1",
            "ca1", com.azure.core.util.Context.NONE);
    }
}
