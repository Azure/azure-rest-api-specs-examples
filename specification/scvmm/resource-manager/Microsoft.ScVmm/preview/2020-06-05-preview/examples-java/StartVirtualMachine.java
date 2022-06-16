import com.azure.core.util.Context;

/** Samples for VirtualMachines Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/StartVirtualMachine.json
     */
    /**
     * Sample code: StartVirtualMachine.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void startVirtualMachine(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachines().start("testrg", "DemoVM", Context.NONE);
    }
}
