
import com.azure.resourcemanager.connectedvmware.models.HardwareProfile;
import com.azure.resourcemanager.connectedvmware.models.VirtualMachineInstanceUpdate;

/**
 * Samples for VirtualMachineInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/
     * UpdateVirtualMachineInstance.json
     */
    /**
     * Sample code: UpdateVirtualMachine.
     * 
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void updateVirtualMachine(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachineInstances().update(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM",
            new VirtualMachineInstanceUpdate().withHardwareProfile(
                new HardwareProfile().withMemorySizeMB(4196).withNumCPUs(4)),
            com.azure.core.util.Context.NONE);
    }
}
