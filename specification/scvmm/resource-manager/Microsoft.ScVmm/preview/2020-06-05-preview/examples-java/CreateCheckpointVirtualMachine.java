import com.azure.core.util.Context;
import com.azure.resourcemanager.scvmm.models.VirtualMachineCreateCheckpoint;

/** Samples for VirtualMachines CreateCheckpoint. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateCheckpointVirtualMachine.json
     */
    /**
     * Sample code: CreateCheckpointVirtualMachine.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void createCheckpointVirtualMachine(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .virtualMachines()
            .createCheckpoint(
                "testrg",
                "DemoVM",
                new VirtualMachineCreateCheckpoint()
                    .withName("Demo Checkpoint name")
                    .withDescription("Demo Checkpoint description"),
                Context.NONE);
    }
}
