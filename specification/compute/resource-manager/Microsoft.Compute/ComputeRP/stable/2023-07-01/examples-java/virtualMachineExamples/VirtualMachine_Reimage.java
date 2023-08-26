import com.azure.resourcemanager.compute.models.VirtualMachineReimageParameters;

/** Samples for VirtualMachines Reimage. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineExamples/VirtualMachine_Reimage.json
     */
    /**
     * Sample code: Reimage a Virtual Machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reimageAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .reimage(
                "myResourceGroup",
                "myVMName",
                new VirtualMachineReimageParameters().withTempDisk(true),
                com.azure.core.util.Context.NONE);
    }
}
