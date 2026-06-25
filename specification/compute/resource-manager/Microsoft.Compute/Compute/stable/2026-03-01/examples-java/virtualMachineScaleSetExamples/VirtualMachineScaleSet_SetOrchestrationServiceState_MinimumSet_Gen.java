
import com.azure.resourcemanager.compute.models.OrchestrationServiceNames;
import com.azure.resourcemanager.compute.models.OrchestrationServiceStateAction;
import com.azure.resourcemanager.compute.models.OrchestrationServiceStateInput;

/**
 * Samples for VirtualMachineScaleSets SetOrchestrationServiceState.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_SetOrchestrationServiceState_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_SetOrchestrationServiceState_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetSetOrchestrationServiceStateMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().setOrchestrationServiceState("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaa",
            new OrchestrationServiceStateInput().withServiceName(OrchestrationServiceNames.AUTOMATIC_REPAIRS)
                .withAction(OrchestrationServiceStateAction.RESUME),
            com.azure.core.util.Context.NONE);
    }
}
