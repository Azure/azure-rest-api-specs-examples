import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetExtensionInner;

/** Samples for VirtualMachineScaleSetExtensions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtensions_CreateOrUpdate_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetExtensionsCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetExtensions()
            .createOrUpdate(
                "rgcompute", "aaaaaaaaaaa", "aaaaaaaaaaa", new VirtualMachineScaleSetExtensionInner(), Context.NONE);
    }
}
