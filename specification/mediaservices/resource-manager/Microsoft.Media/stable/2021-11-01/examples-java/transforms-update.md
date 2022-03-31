Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.BuiltInStandardEncoderPreset;
import com.azure.resourcemanager.mediaservices.models.EncoderNamedPreset;
import com.azure.resourcemanager.mediaservices.models.Priority;
import com.azure.resourcemanager.mediaservices.models.Transform;
import com.azure.resourcemanager.mediaservices.models.TransformOutput;
import java.util.Arrays;

/** Samples for Transforms Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-update.json
     */
    /**
     * Sample code: Update a Transform.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateATransform(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        Transform resource =
            manager
                .transforms()
                .getWithResponse("contosoresources", "contosomedia", "transformToUpdate", Context.NONE)
                .getValue();
        resource
            .update()
            .withDescription("Example transform to illustrate update.")
            .withOutputs(
                Arrays
                    .asList(
                        new TransformOutput()
                            .withRelativePriority(Priority.HIGH)
                            .withPreset(
                                new BuiltInStandardEncoderPreset()
                                    .withPresetName(EncoderNamedPreset.H264MULTIPLE_BITRATE720P))))
            .apply();
    }
}
```
