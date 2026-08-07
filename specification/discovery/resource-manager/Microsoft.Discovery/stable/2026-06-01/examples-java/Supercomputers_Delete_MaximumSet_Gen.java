
/**
 * Samples for Supercomputers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Supercomputers_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Supercomputers_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void supercomputersDeleteMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.supercomputers().delete("rgdiscovery", "7d52dbbe848ddb02a1", com.azure.core.util.Context.NONE);
    }
}
