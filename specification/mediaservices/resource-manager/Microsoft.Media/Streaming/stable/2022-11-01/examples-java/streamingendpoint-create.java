
import com.azure.resourcemanager.mediaservices.models.AkamaiAccessControl;
import com.azure.resourcemanager.mediaservices.models.AkamaiSignatureHeaderAuthenticationKey;
import com.azure.resourcemanager.mediaservices.models.IpAccessControl;
import com.azure.resourcemanager.mediaservices.models.IpRange;
import com.azure.resourcemanager.mediaservices.models.StreamingEndpointAccessControl;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StreamingEndpoints Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/
     * streamingendpoint-create.json
     */
    /**
     * Sample code: Create a streaming endpoint.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().define("myStreamingEndpoint1").withRegion("West US")
            .withExistingMediaservice("mediaresources", "slitestmedia10")
            .withTags(mapOf("tag1", "value1", "tag2", "value2")).withDescription("test event 1").withScaleUnits(1)
            .withAvailabilitySetName("availableset")
            .withAccessControl(new StreamingEndpointAccessControl()
                .withAkamai(new AkamaiAccessControl().withAkamaiSignatureHeaderAuthenticationKeyList(Arrays.asList(
                    new AkamaiSignatureHeaderAuthenticationKey().withIdentifier("id1")
                        .withBase64Key("fakeTokenPlaceholder")
                        .withExpiration(OffsetDateTime.parse("2029-12-31T16:00:00-08:00")),
                    new AkamaiSignatureHeaderAuthenticationKey().withIdentifier("id2")
                        .withBase64Key("fakeTokenPlaceholder")
                        .withExpiration(OffsetDateTime.parse("2030-12-31T16:00:00-08:00")))))
                .withIp(new IpAccessControl()
                    .withAllow(Arrays.asList(new IpRange().withName("AllowedIp").withAddress("192.168.1.1")))))
            .withCdnEnabled(false).create();
    }

    // Use "Map.of" if available
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
