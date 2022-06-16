import com.azure.core.util.Context;

/** Samples for VirtualMachines Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/RestartVirtualMachine.json
     */
    /**
     * Sample code: RestartVirtualMachine.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void restartVirtualMachine(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachines().restart("testrg", "DemoVM", Context.NONE);
    }
}
