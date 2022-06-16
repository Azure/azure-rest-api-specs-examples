import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetExtensionUpdate;

/** Samples for VirtualMachineScaleSetExtensions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtensions_Update_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetExtensionsUpdateMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetExtensions()
            .update(
                "rgcompute",
                "aaaaaaaaaaaaaaaaaaaaaaaaaa",
                "aa",
                new VirtualMachineScaleSetExtensionUpdate(),
                Context.NONE);
    }
}
