import com.azure.core.util.Context;

/** Samples for DiskEncryptionSets ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_ListByResourceGroup.json
     */
    /**
     * Sample code: List all disk encryption sets in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllDiskEncryptionSetsInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskEncryptionSets()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
