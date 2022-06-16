import com.azure.resourcemanager.mediaservices.models.BuiltInStandardEncoderPreset;
import com.azure.resourcemanager.mediaservices.models.EncoderNamedPreset;
import com.azure.resourcemanager.mediaservices.models.TransformOutput;
import java.util.Arrays;

/** Samples for Transforms CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-create.json
     */
    /**
     * Sample code: Create or update a Transform.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createOrUpdateATransform(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .transforms()
            .define("createdTransform")
            .withExistingMediaService("contosoresources", "contosomedia")
            .withDescription("Example Transform to illustrate create and update.")
            .withOutputs(
                Arrays
                    .asList(
                        new TransformOutput()
                            .withPreset(
                                new BuiltInStandardEncoderPreset()
                                    .withPresetName(EncoderNamedPreset.ADAPTIVE_STREAMING))))
            .create();
    }
}
