import com.azure.core.util.Context;

/** Samples for Location ListCachedImages. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/CachedImagesList.json
     */
    /**
     * Sample code: CachedImages.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cachedImages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerGroups()
            .manager()
            .serviceClient()
            .getLocations()
            .listCachedImages("westcentralus", Context.NONE);
    }
}
