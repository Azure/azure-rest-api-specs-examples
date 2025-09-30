
import com.azure.resourcemanager.compute.models.VMScaleSetScaleOutInput;
import com.azure.resourcemanager.compute.models.VMScaleSetScaleOutInputProperties;

/**
 * Samples for VirtualMachineScaleSets ScaleOut.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_ScaleOut.json
     */
    /**
     * Sample code: VirtualMachineScaleSet Scale Out.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetScaleOut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().scaleOut("myResourceGroup",
            "{vmss-name}", new VMScaleSetScaleOutInput().withCapacity(5L).withProperties(
                new VMScaleSetScaleOutInputProperties().withZone("1")),
            com.azure.core.util.Context.NONE);
    }
}
