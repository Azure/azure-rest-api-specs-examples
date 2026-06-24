
import com.azure.resourcemanager.compute.models.InterconnectBlockExpandTypes;

/**
 * Samples for InterconnectBlocks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_Get_InstanceView.json
     */
    /**
     * Sample code: Get an Interconnect Block with instance view.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAnInterconnectBlockWithInstanceView(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getInterconnectBlocks().getByResourceGroupWithResponse("myResourceGroup",
            "myInterconnectBlock", InterconnectBlockExpandTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
