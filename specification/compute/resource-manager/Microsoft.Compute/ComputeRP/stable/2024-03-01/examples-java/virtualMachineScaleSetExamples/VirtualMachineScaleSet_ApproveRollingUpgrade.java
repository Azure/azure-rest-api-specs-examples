
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceIDs;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets ApproveRollingUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_ApproveRollingUpgrade.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ApproveRollingUpgrade.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetApproveRollingUpgrade(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().approveRollingUpgrade(
            "rgcompute", "vmssToApproveRollingUpgradeOn",
            new VirtualMachineScaleSetVMInstanceIDs().withInstanceIds(Arrays.asList("0", "1", "2")),
            com.azure.core.util.Context.NONE);
    }
}
