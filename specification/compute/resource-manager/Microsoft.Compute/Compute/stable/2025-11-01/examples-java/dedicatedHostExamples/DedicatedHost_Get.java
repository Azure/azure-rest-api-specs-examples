
import com.azure.resourcemanager.compute.models.InstanceViewTypes;

/**
 * Samples for DedicatedHosts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHost_Get.json
     */
    /**
     * Sample code: Get a dedicated host.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getADedicatedHost(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHosts().getWithResponse("myResourceGroup", "myDedicatedHostGroup", "myHost",
            InstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
