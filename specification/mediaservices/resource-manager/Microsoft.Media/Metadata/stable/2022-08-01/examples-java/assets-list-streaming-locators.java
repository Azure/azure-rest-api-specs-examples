
/**
 * Samples for Assets ListStreamingLocators.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-list-
     * streaming-locators.json
     */
    /**
     * Sample code: List Asset SAS URLs.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAssetSASURLs(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.assets().listStreamingLocatorsWithResponse("contosorg", "contosomedia", "ClimbingMountSaintHelens",
            com.azure.core.util.Context.NONE);
    }
}
