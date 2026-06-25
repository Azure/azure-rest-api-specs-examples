
/**
 * Samples for InterconnectBlocks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_ListByResourceGroup.json
     */
    /**
     * Sample code: List Interconnect Blocks in a resource group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listInterconnectBlocksInAResourceGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getInterconnectBlocks().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
