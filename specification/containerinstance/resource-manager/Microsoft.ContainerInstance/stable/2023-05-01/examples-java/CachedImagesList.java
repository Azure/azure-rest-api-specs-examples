/** Samples for Location ListCachedImages. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/CachedImagesList.json
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
            .listCachedImages("westcentralus", com.azure.core.util.Context.NONE);
    }
}
