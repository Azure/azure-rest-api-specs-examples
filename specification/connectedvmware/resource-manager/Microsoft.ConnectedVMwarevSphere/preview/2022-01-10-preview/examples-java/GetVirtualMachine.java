import com.azure.core.util.Context;

/** Samples for VirtualMachines GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetVirtualMachine.json
     */
    /**
     * Sample code: GetVirtualMachine.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getVirtualMachine(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachines().getByResourceGroupWithResponse("testrg", "DemoVM", Context.NONE);
    }
}
