```java
import com.azure.resourcemanager.mediaservices.models.CbcsDrmConfiguration;
import com.azure.resourcemanager.mediaservices.models.CommonEncryptionCbcs;
import com.azure.resourcemanager.mediaservices.models.DefaultKey;
import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyContentKeys;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyFairPlayConfiguration;

/** Samples for StreamingPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-policies-create-commonEncryptionCbcs-only.json
     */
    /**
     * Sample code: Creates a Streaming Policy with commonEncryptionCbcs only.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithCommonEncryptionCbcsOnly(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingPolicies()
            .define("UserCreatedSecureStreamingPolicyWithCommonEncryptionCbcsOnly")
            .withExistingMediaService("contoso", "contosomedia")
            .withDefaultContentKeyPolicyName("PolicyWithMultipleOptions")
            .withCommonEncryptionCbcs(
                new CommonEncryptionCbcs()
                    .withEnabledProtocols(
                        new EnabledProtocols()
                            .withDownload(false)
                            .withDash(false)
                            .withHls(true)
                            .withSmoothStreaming(false))
                    .withContentKeys(
                        new StreamingPolicyContentKeys().withDefaultKey(new DefaultKey().withLabel("cbcsDefaultKey")))
                    .withDrm(
                        new CbcsDrmConfiguration()
                            .withFairPlay(
                                new StreamingPolicyFairPlayConfiguration()
                                    .withCustomLicenseAcquisitionUrlTemplate(
                                        "https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}")
                                    .withAllowPersistentLicense(true))))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.
