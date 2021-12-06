Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.4/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.videoanalyzer.models.Kind;
import com.azure.resourcemanager.videoanalyzer.models.NodeInput;
import com.azure.resourcemanager.videoanalyzer.models.ParameterDeclaration;
import com.azure.resourcemanager.videoanalyzer.models.ParameterType;
import com.azure.resourcemanager.videoanalyzer.models.RtspSource;
import com.azure.resourcemanager.videoanalyzer.models.RtspTransport;
import com.azure.resourcemanager.videoanalyzer.models.Sku;
import com.azure.resourcemanager.videoanalyzer.models.SkuName;
import com.azure.resourcemanager.videoanalyzer.models.UnsecuredEndpoint;
import com.azure.resourcemanager.videoanalyzer.models.UsernamePasswordCredentials;
import com.azure.resourcemanager.videoanalyzer.models.VideoCreationProperties;
import com.azure.resourcemanager.videoanalyzer.models.VideoPublishingOptions;
import com.azure.resourcemanager.videoanalyzer.models.VideoSink;
import java.util.Arrays;

/** Samples for PipelineTopologies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-create.json
     */
    /**
     * Sample code: Create or update a pipeline topology with an Rtsp source and video sink.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void createOrUpdateAPipelineTopologyWithAnRtspSourceAndVideoSink(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .pipelineTopologies()
            .define("pipelineTopology1")
            .withExistingVideoAnalyzer("testrg", "testaccount2")
            .withKind(Kind.LIVE)
            .withSku(new Sku().withName(SkuName.LIVE_S1))
            .withDescription("Pipeline Topology 1 Description")
            .withParameters(
                Arrays
                    .asList(
                        new ParameterDeclaration()
                            .withName("rtspUrlParameter")
                            .withType(ParameterType.STRING)
                            .withDescription("rtsp source url parameter")
                            .withDefaultProperty("rtsp://microsoft.com/video.mp4"),
                        new ParameterDeclaration()
                            .withName("rtspPasswordParameter")
                            .withType(ParameterType.SECRET_STRING)
                            .withDescription("rtsp source password parameter")
                            .withDefaultProperty("password")))
            .withSources(
                Arrays
                    .asList(
                        new RtspSource()
                            .withName("rtspSource")
                            .withTransport(RtspTransport.HTTP)
                            .withEndpoint(
                                new UnsecuredEndpoint()
                                    .withCredentials(
                                        new UsernamePasswordCredentials()
                                            .withUsername("username")
                                            .withPassword("${rtspPasswordParameter}"))
                                    .withUrl("${rtspUrlParameter}"))))
            .withSinks(
                Arrays
                    .asList(
                        new VideoSink()
                            .withName("videoSink")
                            .withInputs(Arrays.asList(new NodeInput().withNodeName("rtspSource")))
                            .withVideoName("camera001")
                            .withVideoCreationProperties(
                                new VideoCreationProperties()
                                    .withTitle("Parking Lot (Camera 1)")
                                    .withDescription("Parking lot south entrance")
                                    .withSegmentLength("PT30S"))
                            .withVideoPublishingOptions(
                                new VideoPublishingOptions()
                                    .withDisableArchive("false")
                                    .withDisableRtspPublishing("true"))))
            .create();
    }
}
```
