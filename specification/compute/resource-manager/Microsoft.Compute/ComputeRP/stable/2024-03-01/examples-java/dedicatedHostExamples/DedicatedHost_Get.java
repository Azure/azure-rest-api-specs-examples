
import com.azure.resourcemanager.compute.models.InstanceViewTypes;

/**
 * Samples for DedicatedHosts Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * dedicatedHostExamples/DedicatedHost_Get.json
     */
    /**
     * Sample code: Get a dedicated host.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADedicatedHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHosts().getWithResponse("myResourceGroup",
            "myDedicatedHostGroup", "myHost", InstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
