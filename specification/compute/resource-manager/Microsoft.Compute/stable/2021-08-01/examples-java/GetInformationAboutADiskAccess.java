import com.azure.core.util.Context;

/** Samples for DiskAccesses GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/GetInformationAboutADiskAccess.json
     */
    /**
     * Sample code: Get information about a disk access resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutADiskAccessResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .getByResourceGroupWithResponse("myResourceGroup", "myDiskAccess", Context.NONE);
    }
}
