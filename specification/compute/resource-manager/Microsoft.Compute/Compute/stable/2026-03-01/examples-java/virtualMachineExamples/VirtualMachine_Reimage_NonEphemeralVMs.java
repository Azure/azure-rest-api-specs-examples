
import com.azure.resourcemanager.compute.models.OSProfileProvisioningData;
import com.azure.resourcemanager.compute.models.VirtualMachineReimageParameters;

/**
 * Samples for VirtualMachines Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Reimage_NonEphemeralVMs.json
     */
    /**
     * Sample code: Reimage a Non-Ephemeral Virtual Machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void reimageANonEphemeralVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().reimage(
            "myResourceGroup", "myVMName", new VirtualMachineReimageParameters().withTempDisk(true)
                .withExactVersion("aaaaaa").withOsProfile(new OSProfileProvisioningData()
                    .withAdminPassword("fakeTokenPlaceholder").withCustomData("{your-custom-data}")),
            com.azure.core.util.Context.NONE);
    }
}
