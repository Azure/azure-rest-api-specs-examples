import com.azure.core.util.Context;

/** Samples for Videos Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-get.json
     */
    /**
     * Sample code: Gets a video entity.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getsAVideoEntity(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.videos().getWithResponse("testrg", "testaccount2", "video1", Context.NONE);
    }
}
