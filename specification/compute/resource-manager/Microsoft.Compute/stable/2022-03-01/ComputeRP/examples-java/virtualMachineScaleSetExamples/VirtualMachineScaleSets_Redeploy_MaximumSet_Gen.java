import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceIDs;
import java.util.Arrays;

/** Samples for VirtualMachineScaleSets Redeploy. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_Redeploy_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_Redeploy_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsRedeployMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .redeploy(
                "rgcompute",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                new VirtualMachineScaleSetVMInstanceIDs().withInstanceIds(Arrays.asList("aaaaaaaaaaaaaaaaa")),
                Context.NONE);
    }
}
