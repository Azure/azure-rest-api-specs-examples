import com.azure.core.util.Context;

/** Samples for VirtualMachines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVirtualMachinesByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualMachinesByResourceGroup.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVirtualMachinesByResourceGroup(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachines().listByResourceGroup("testrg", Context.NONE);
    }
}
