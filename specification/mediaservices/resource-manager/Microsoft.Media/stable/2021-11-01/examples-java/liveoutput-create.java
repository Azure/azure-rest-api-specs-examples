import com.azure.resourcemanager.mediaservices.models.Hls;
import java.time.Duration;

/** Samples for LiveOutputs Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveoutput-create.json
     */
    /**
     * Sample code: Create a LiveOutput.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createALiveOutput(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveOutputs()
            .define("myLiveOutput1")
            .withExistingLiveEvent("mediaresources", "slitestmedia10", "myLiveEvent1")
            .withDescription("test live output 1")
            .withAssetName("6f3264f5-a189-48b4-a29a-a40f22575212")
            .withArchiveWindowLength(Duration.parse("PT5M"))
            .withManifestName("testmanifest")
            .withHls(new Hls().withFragmentsPerTsSegment(5))
            .create();
    }
}
