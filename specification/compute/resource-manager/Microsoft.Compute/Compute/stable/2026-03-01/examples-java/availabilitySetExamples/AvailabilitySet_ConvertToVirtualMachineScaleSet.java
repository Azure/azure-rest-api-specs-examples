
import com.azure.resourcemanager.compute.models.ConvertToVirtualMachineScaleSetInput;

/**
 * Samples for AvailabilitySets ConvertToVirtualMachineScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_ConvertToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: AvailabilitySet_ConvertToVirtualMachineScaleSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        availabilitySetConvertToVirtualMachineScaleSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().convertToVirtualMachineScaleSet("rgcompute", "myAvailabilitySet",
            new ConvertToVirtualMachineScaleSetInput().withVirtualMachineScaleSetName("{vmss-name}"),
            com.azure.core.util.Context.NONE);
    }
}
