
import com.azure.resourcemanager.compute.models.OrchestrationServiceNames;
import com.azure.resourcemanager.compute.models.OrchestrationServiceStateAction;
import com.azure.resourcemanager.compute.models.OrchestrationServiceStateInput;

/**
 * Samples for VirtualMachineScaleSets SetOrchestrationServiceState.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_SetOrchestrationServiceState_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_SetOrchestrationServiceState_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetSetOrchestrationServiceStateMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().setOrchestrationServiceState(
            "rgcompute", "aaaaaaaaaaaaaaaa",
            new OrchestrationServiceStateInput().withServiceName(OrchestrationServiceNames.AUTOMATIC_REPAIRS)
                .withAction(OrchestrationServiceStateAction.RESUME),
            com.azure.core.util.Context.NONE);
    }
}
