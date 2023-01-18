import com.azure.resourcemanager.mediaservices.models.StreamingLocatorContentKey;
import java.util.Arrays;
import java.util.UUID;

/** Samples for StreamingLocators Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-secure-userDefinedContentKeys.json
     */
    /**
     * Sample code: Creates a Streaming Locator with user defined content keys.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingLocatorWithUserDefinedContentKeys(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingLocators()
            .define("UserCreatedSecureStreamingLocatorWithUserDefinedContentKeys")
            .withExistingMediaService("contoso", "contosomedia")
            .withAssetName("ClimbingMountRainier")
            .withStreamingLocatorId(UUID.fromString("90000000-0000-0000-0000-00000000000A"))
            .withStreamingPolicyName("secureStreamingPolicy")
            .withContentKeys(
                Arrays
                    .asList(
                        new StreamingLocatorContentKey()
                            .withId(UUID.fromString("60000000-0000-0000-0000-000000000001"))
                            .withLabelReferenceInStreamingPolicy("aesDefaultKey")
                            .withValue("1UqLohAfWsEGkULYxHjYZg=="),
                        new StreamingLocatorContentKey()
                            .withId(UUID.fromString("60000000-0000-0000-0000-000000000004"))
                            .withLabelReferenceInStreamingPolicy("cencDefaultKey")
                            .withValue("4UqLohAfWsEGkULYxHjYZg=="),
                        new StreamingLocatorContentKey()
                            .withId(UUID.fromString("60000000-0000-0000-0000-000000000007"))
                            .withLabelReferenceInStreamingPolicy("cbcsDefaultKey")
                            .withValue("7UqLohAfWsEGkULYxHjYZg==")))
            .create();
    }
}
