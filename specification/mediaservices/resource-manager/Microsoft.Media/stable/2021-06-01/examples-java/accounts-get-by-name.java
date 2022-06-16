import com.azure.core.util.Context;

/** Samples for Mediaservices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-get-by-name.json
     */
    /**
     * Sample code: Get a Media Services account by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAMediaServicesAccountByName(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.mediaservices().getByResourceGroupWithResponse("contoso", "contosotv", Context.NONE);
    }
}
