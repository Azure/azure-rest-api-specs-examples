import com.azure.core.util.Context;
import com.azure.resourcemanager.videoanalyzer.models.VideoEntity;

/** Samples for Videos Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-patch.json
     */
    /**
     * Sample code: Update video entity.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void updateVideoEntity(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        VideoEntity resource =
            manager.videos().getWithResponse("testrg", "testaccount2", "video1", Context.NONE).getValue();
        resource.update().withDescription("Parking Lot East Entrance").apply();
    }
}
