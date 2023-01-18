/** Samples for StreamingLocators ListPaths. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list-paths-streaming-only.json
     */
    /**
     * Sample code: List Paths which has streaming paths only.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listPathsWhichHasStreamingPathsOnly(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingLocators()
            .listPathsWithResponse(
                "contoso", "contosomedia", "secureStreamingLocator", com.azure.core.util.Context.NONE);
    }
}
