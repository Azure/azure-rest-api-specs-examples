Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mediaservices.models.CencDrmConfiguration;
import com.azure.resourcemanager.mediaservices.models.CommonEncryptionCenc;
import com.azure.resourcemanager.mediaservices.models.DefaultKey;
import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyContentKeys;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyPlayReadyConfiguration;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyWidevineConfiguration;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyCompareOperation;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyCondition;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyType;
import com.azure.resourcemanager.mediaservices.models.TrackSelection;
import java.util.Arrays;

/** Samples for StreamingPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-policies-create-commonEncryptionCenc-only.json
     */
    /**
     * Sample code: Creates a Streaming Policy with commonEncryptionCenc only.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithCommonEncryptionCencOnly(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingPolicies()
            .define("UserCreatedSecureStreamingPolicyWithCommonEncryptionCencOnly")
            .withExistingMediaService("contoso", "contosomedia")
            .withDefaultContentKeyPolicyName("PolicyWithPlayReadyOptionAndOpenRestriction")
            .withCommonEncryptionCenc(
                new CommonEncryptionCenc()
                    .withEnabledProtocols(
                        new EnabledProtocols()
                            .withDownload(false)
                            .withDash(true)
                            .withHls(false)
                            .withSmoothStreaming(true))
                    .withClearTracks(
                        Arrays
                            .asList(
                                new TrackSelection()
                                    .withTrackSelections(
                                        Arrays
                                            .asList(
                                                new TrackPropertyCondition()
                                                    .withProperty(TrackPropertyType.FOUR_CC)
                                                    .withOperation(TrackPropertyCompareOperation.EQUAL)
                                                    .withValue("hev1")))))
                    .withContentKeys(
                        new StreamingPolicyContentKeys().withDefaultKey(new DefaultKey().withLabel("cencDefaultKey")))
                    .withDrm(
                        new CencDrmConfiguration()
                            .withPlayReady(
                                new StreamingPolicyPlayReadyConfiguration()
                                    .withCustomLicenseAcquisitionUrlTemplate(
                                        "https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}")
                                    .withPlayReadyCustomAttributes("PlayReady CustomAttributes"))
                            .withWidevine(
                                new StreamingPolicyWidevineConfiguration()
                                    .withCustomLicenseAcquisitionUrlTemplate(
                                        "https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId"))))
            .create();
    }
}
```
