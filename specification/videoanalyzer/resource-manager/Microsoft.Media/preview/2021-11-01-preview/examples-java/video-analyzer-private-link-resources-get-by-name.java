import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-private-link-resources-get-by-name.json
     */
    /**
     * Sample code: Get details of a group ID.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getDetailsOfAGroupID(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.privateLinkResources().getWithResponse("contoso", "contososports", "integration", Context.NONE);
    }
}
