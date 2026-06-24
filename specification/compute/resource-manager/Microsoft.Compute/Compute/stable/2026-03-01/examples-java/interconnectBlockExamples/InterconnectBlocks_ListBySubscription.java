
/**
 * Samples for InterconnectBlocks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_ListBySubscription.json
     */
    /**
     * Sample code: List Interconnect Blocks in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listInterconnectBlocksInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getInterconnectBlocks().list(com.azure.core.util.Context.NONE);
    }
}
