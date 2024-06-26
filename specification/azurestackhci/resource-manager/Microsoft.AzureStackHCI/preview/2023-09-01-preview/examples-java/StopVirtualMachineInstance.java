/** Samples for VirtualMachineInstances Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/StopVirtualMachineInstance.json
     */
    /**
     * Sample code: StopVirtualMachine.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void stopVirtualMachine(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .virtualMachineInstances()
            .stop(
                "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default",
                com.azure.core.util.Context.NONE);
    }
}
