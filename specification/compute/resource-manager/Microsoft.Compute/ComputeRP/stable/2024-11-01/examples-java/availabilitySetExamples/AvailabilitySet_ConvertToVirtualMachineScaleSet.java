
import com.azure.resourcemanager.compute.models.ConvertToVirtualMachineScaleSetInput;

/**
 * Samples for AvailabilitySets ConvertToVirtualMachineScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * availabilitySetExamples/AvailabilitySet_ConvertToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: AvailabilitySet_ConvertToVirtualMachineScaleSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        availabilitySetConvertToVirtualMachineScaleSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getAvailabilitySets().convertToVirtualMachineScaleSet(
            "rgcompute", "myAvailabilitySet",
            new ConvertToVirtualMachineScaleSetInput().withVirtualMachineScaleSetName("{vmss-name}"),
            com.azure.core.util.Context.NONE);
    }
}
