/** Samples for VirtualMachineInstances Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/StartVirtualMachineInstance.json
     */
    /**
     * Sample code: StartVirtualMachine.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void startVirtualMachine(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .virtualMachineInstances()
            .start(
                "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM",
                com.azure.core.util.Context.NONE);
    }
}
