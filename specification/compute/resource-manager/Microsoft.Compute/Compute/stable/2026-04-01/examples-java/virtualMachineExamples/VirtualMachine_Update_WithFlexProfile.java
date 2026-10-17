
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineUpdateInner;
import com.azure.resourcemanager.compute.models.FlexProfile;

/**
 * Samples for VirtualMachines Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/virtualMachineExamples/VirtualMachine_Update_WithFlexProfile.json
     */
    /**
     * Sample code: Update a VM's FlexProfile.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateAVMSFlexProfile(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().update("myResourceGroup", "myVM",
            new VirtualMachineUpdateInner().withFlexProfile(new FlexProfile().withPreferredVMSize("Standard_D8s_v5")),
            null, null, com.azure.core.util.Context.NONE);
    }
}
