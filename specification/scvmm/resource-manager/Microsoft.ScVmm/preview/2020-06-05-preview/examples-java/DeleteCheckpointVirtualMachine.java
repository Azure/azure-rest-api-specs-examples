import com.azure.core.util.Context;
import com.azure.resourcemanager.scvmm.models.VirtualMachineDeleteCheckpoint;

/** Samples for VirtualMachines DeleteCheckpoint. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteCheckpointVirtualMachine.json
     */
    /**
     * Sample code: DeleteCheckpointVirtualMachine.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void deleteCheckpointVirtualMachine(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .virtualMachines()
            .deleteCheckpoint(
                "testrg", "DemoVM", new VirtualMachineDeleteCheckpoint().withId("Demo CheckpointID"), Context.NONE);
    }
}
