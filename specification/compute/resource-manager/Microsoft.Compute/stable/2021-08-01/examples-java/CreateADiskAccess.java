import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.DiskAccessInner;

/** Samples for DiskAccesses CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/CreateADiskAccess.json
     */
    /**
     * Sample code: Create a disk access resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADiskAccessResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .createOrUpdate(
                "myResourceGroup", "myDiskAccess", new DiskAccessInner().withLocation("West US"), Context.NONE);
    }
}
