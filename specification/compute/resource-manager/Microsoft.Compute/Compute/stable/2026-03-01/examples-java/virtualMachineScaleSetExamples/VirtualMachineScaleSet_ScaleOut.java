
import com.azure.resourcemanager.compute.models.VMScaleSetScaleOutInput;
import com.azure.resourcemanager.compute.models.VMScaleSetScaleOutInputProperties;

/**
 * Samples for VirtualMachineScaleSets ScaleOut.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ScaleOut.json
     */
    /**
     * Sample code: VirtualMachineScaleSet Scale Out.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetScaleOut(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().scaleOut("myResourceGroup", "{vmss-name}",
            new VMScaleSetScaleOutInput().withCapacity(5L).withProperties(
                new VMScaleSetScaleOutInputProperties().withZone("1")),
            com.azure.core.util.Context.NONE);
    }
}
