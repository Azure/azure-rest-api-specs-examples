import com.azure.core.util.Context;

/** Samples for DiskEncryptionSets ListAssociatedResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/ListDiskEncryptionSetAssociatedResources.json
     */
    /**
     * Sample code: List all resources that are encrypted with this disk encryption set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllResourcesThatAreEncryptedWithThisDiskEncryptionSet(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskEncryptionSets()
            .listAssociatedResources("myResourceGroup", "myDiskEncryptionSet", Context.NONE);
    }
}
