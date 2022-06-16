import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.IpAccessControl;
import com.azure.resourcemanager.mediaservices.models.IpRange;
import com.azure.resourcemanager.mediaservices.models.LiveEvent;
import com.azure.resourcemanager.mediaservices.models.LiveEventInput;
import com.azure.resourcemanager.mediaservices.models.LiveEventInputAccessControl;
import com.azure.resourcemanager.mediaservices.models.LiveEventInputProtocol;
import com.azure.resourcemanager.mediaservices.models.LiveEventPreview;
import com.azure.resourcemanager.mediaservices.models.LiveEventPreviewAccessControl;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for LiveEvents Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-update.json
     */
    /**
     * Sample code: Update a LiveEvent.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateALiveEvent(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        LiveEvent resource =
            manager
                .liveEvents()
                .getWithResponse("mediaresources", "slitestmedia10", "myLiveEvent1", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("tag1", "value1", "tag2", "value2", "tag3", "value3"))
            .withDescription("test event updated")
            .withInput(
                new LiveEventInput()
                    .withStreamingProtocol(LiveEventInputProtocol.FRAGMENTED_MP4)
                    .withAccessControl(
                        new LiveEventInputAccessControl()
                            .withIp(
                                new IpAccessControl()
                                    .withAllow(
                                        Arrays.asList(new IpRange().withName("AllowOne").withAddress("192.1.1.0")))))
                    .withKeyFrameIntervalDuration("PT6S"))
            .withPreview(
                new LiveEventPreview()
                    .withAccessControl(
                        new LiveEventPreviewAccessControl()
                            .withIp(
                                new IpAccessControl()
                                    .withAllow(
                                        Arrays.asList(new IpRange().withName("AllowOne").withAddress("192.1.1.0"))))))
            .apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
