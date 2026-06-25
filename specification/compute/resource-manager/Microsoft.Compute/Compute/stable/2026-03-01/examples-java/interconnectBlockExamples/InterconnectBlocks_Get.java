
/**
 * Samples for InterconnectBlocks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_Get.json
     */
    /**
     * Sample code: Get an Interconnect Block.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAnInterconnectBlock(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getInterconnectBlocks().getByResourceGroupWithResponse("myResourceGroup",
            "myInterconnectBlock", null, com.azure.core.util.Context.NONE);
    }
}
