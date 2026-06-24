
/**
 * Samples for InterconnectBlocks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_Delete.json
     */
    /**
     * Sample code: Delete an Interconnect Block.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAnInterconnectBlock(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getInterconnectBlocks().delete("myResourceGroup", "myInterconnectBlock",
            com.azure.core.util.Context.NONE);
    }
}
