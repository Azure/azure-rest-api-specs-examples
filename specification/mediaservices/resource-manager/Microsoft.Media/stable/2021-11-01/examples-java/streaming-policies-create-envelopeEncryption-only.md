Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mediaservices.models.DefaultKey;
import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.EnvelopeEncryption;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyContentKeys;

/** Samples for StreamingPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-policies-create-envelopeEncryption-only.json
     */
    /**
     * Sample code: Creates a Streaming Policy with envelopeEncryption only.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithEnvelopeEncryptionOnly(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingPolicies()
            .define("UserCreatedSecureStreamingPolicyWithEnvelopeEncryptionOnly")
            .withExistingMediaService("contoso", "contosomedia")
            .withDefaultContentKeyPolicyName("PolicyWithClearKeyOptionAndTokenRestriction")
            .withEnvelopeEncryption(
                new EnvelopeEncryption()
                    .withEnabledProtocols(
                        new EnabledProtocols()
                            .withDownload(false)
                            .withDash(true)
                            .withHls(true)
                            .withSmoothStreaming(true))
                    .withContentKeys(
                        new StreamingPolicyContentKeys().withDefaultKey(new DefaultKey().withLabel("aesDefaultKey")))
                    .withCustomKeyAcquisitionUrlTemplate(
                        "https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}"))
            .create();
    }
}
```
