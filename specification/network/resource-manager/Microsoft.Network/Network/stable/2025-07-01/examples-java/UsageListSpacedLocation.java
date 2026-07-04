
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/UsageListSpacedLocation.json
     */
    /**
     * Sample code: List usages spaced location.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listUsagesSpacedLocation(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getUsages().list("West US", com.azure.core.util.Context.NONE);
    }
}
