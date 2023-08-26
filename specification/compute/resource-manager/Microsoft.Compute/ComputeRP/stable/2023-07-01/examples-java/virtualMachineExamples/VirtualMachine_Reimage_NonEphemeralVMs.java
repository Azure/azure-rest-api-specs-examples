import com.azure.resourcemanager.compute.models.OSProfileProvisioningData;
import com.azure.resourcemanager.compute.models.VirtualMachineReimageParameters;

/** Samples for VirtualMachines Reimage. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineExamples/VirtualMachine_Reimage_NonEphemeralVMs.json
     */
    /**
     * Sample code: Reimage a Non-Ephemeral Virtual Machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reimageANonEphemeralVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .reimage(
                "myResourceGroup",
                "myVMName",
                new VirtualMachineReimageParameters()
                    .withTempDisk(true)
                    .withExactVersion("aaaaaa")
                    .withOsProfile(
                        new OSProfileProvisioningData()
                            .withAdminPassword("fakeTokenPlaceholder")
                            .withCustomData("{your-custom-data}")),
                com.azure.core.util.Context.NONE);
    }
}
