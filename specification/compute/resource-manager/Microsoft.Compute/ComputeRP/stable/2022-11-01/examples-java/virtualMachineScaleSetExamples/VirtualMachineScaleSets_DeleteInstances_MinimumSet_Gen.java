import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceRequiredIDs;
import java.util.Arrays;

/** Samples for VirtualMachineScaleSets DeleteInstances. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_DeleteInstances_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_DeleteInstances_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsDeleteInstancesMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .deleteInstances(
                "rgcompute",
                "aaaaaaaaaaaaaaa",
                new VirtualMachineScaleSetVMInstanceRequiredIDs()
                    .withInstanceIds(Arrays.asList("aaaaaaaaaaaaaaaaaaaaaaaaa")),
                null,
                com.azure.core.util.Context.NONE);
    }
}
