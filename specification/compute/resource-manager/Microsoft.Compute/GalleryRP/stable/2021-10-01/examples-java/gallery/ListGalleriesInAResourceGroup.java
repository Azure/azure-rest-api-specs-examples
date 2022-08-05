import com.azure.core.util.Context;

/** Samples for Galleries ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/ListGalleriesInAResourceGroup.json
     */
    /**
     * Sample code: List galleries in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleriesInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleries()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
