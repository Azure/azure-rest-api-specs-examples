/** Samples for HybridIdentityMetadata List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/HybridIdentityMetadata_List.json
     */
    /**
     * Sample code: HybridIdentityMetadataListByVirtualMachineInstances.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void hybridIdentityMetadataListByVirtualMachineInstances(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .hybridIdentityMetadatas()
            .list(
                "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM",
                com.azure.core.util.Context.NONE);
    }
}
