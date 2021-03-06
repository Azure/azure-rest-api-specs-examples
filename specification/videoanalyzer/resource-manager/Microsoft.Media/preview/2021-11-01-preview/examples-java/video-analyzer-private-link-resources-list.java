import com.azure.core.util.Context;

/** Samples for PrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-private-link-resources-list.json
     */
    /**
     * Sample code: Get list of all group IDs.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getListOfAllGroupIDs(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.privateLinkResources().listWithResponse("contoso", "contososports", Context.NONE);
    }
}
