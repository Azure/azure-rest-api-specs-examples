/** Samples for Mediaservices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/accounts-get-by-name.json
     */
    /**
     * Sample code: Get a Media Services account by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAMediaServicesAccountByName(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .mediaservices()
            .getByResourceGroupWithResponse("contosorg", "contosotv", com.azure.core.util.Context.NONE);
    }
}
