
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceIDs;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets ApproveRollingUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ApproveRollingUpgrade.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ApproveRollingUpgrade.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetApproveRollingUpgrade(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().approveRollingUpgrade("rgcompute",
            "vmssToApproveRollingUpgradeOn",
            new VirtualMachineScaleSetVMInstanceIDs().withInstanceIds(Arrays.asList("0", "1", "2")),
            com.azure.core.util.Context.NONE);
    }
}
