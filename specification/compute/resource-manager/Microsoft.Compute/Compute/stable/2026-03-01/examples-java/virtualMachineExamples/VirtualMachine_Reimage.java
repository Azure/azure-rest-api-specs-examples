
import com.azure.resourcemanager.compute.models.VirtualMachineReimageParameters;

/**
 * Samples for VirtualMachines Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Reimage.json
     */
    /**
     * Sample code: Reimage a Virtual Machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void reimageAVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().reimage("myResourceGroup", "myVMName",
            new VirtualMachineReimageParameters().withTempDisk(true), com.azure.core.util.Context.NONE);
    }
}
