import com.azure.core.util.Context;

/** Samples for VirtualMachines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteVirtualMachine.json
     */
    /**
     * Sample code: DeleteVirtualMachine.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void deleteVirtualMachine(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachines().delete("testrg", "DemoVM", null, null, Context.NONE);
    }
}
