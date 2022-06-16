import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceIDs;
import java.util.Arrays;

/** Samples for VirtualMachineScaleSets Deallocate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Deallocate_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_Deallocate_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsDeallocateMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .deallocate(
                "rgcompute",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                new VirtualMachineScaleSetVMInstanceIDs().withInstanceIds(Arrays.asList("aaaaaaaaaaaaaaaaa")),
                Context.NONE);
    }
}
