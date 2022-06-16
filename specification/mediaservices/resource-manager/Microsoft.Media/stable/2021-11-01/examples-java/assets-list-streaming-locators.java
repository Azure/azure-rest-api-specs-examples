import com.azure.core.util.Context;

/** Samples for Assets ListStreamingLocators. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-list-streaming-locators.json
     */
    /**
     * Sample code: List Asset SAS URLs.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAssetSASURLs(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assets()
            .listStreamingLocatorsWithResponse("contoso", "contosomedia", "ClimbingMountSaintHelens", Context.NONE);
    }
}
