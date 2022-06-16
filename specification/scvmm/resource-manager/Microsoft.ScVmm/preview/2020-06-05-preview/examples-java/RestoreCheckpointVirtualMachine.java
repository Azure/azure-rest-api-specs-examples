import com.azure.core.util.Context;
import com.azure.resourcemanager.scvmm.models.VirtualMachineRestoreCheckpoint;

/** Samples for VirtualMachines RestoreCheckpoint. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/RestoreCheckpointVirtualMachine.json
     */
    /**
     * Sample code: RestoreCheckpointVirtualMachine.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void restoreCheckpointVirtualMachine(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .virtualMachines()
            .restoreCheckpoint(
                "testrg", "DemoVM", new VirtualMachineRestoreCheckpoint().withId("Demo CheckpointID"), Context.NONE);
    }
}
