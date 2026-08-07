
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void operationsListMinimumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
