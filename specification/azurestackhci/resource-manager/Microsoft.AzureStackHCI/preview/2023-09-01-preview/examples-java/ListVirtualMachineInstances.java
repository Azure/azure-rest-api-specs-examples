/** Samples for VirtualMachineInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListVirtualMachineInstances.json
     */
    /**
     * Sample code: ListVirtualMachineInstances.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listVirtualMachineInstances(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .virtualMachineInstances()
            .list(
                "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM",
                com.azure.core.util.Context.NONE);
    }
}
