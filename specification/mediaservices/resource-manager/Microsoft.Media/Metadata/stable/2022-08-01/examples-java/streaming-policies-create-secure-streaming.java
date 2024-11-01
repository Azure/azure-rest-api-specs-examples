
import com.azure.resourcemanager.mediaservices.models.CbcsDrmConfiguration;
import com.azure.resourcemanager.mediaservices.models.CencDrmConfiguration;
import com.azure.resourcemanager.mediaservices.models.CommonEncryptionCbcs;
import com.azure.resourcemanager.mediaservices.models.CommonEncryptionCenc;
import com.azure.resourcemanager.mediaservices.models.DefaultKey;
import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.EnvelopeEncryption;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyContentKeys;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyFairPlayConfiguration;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyPlayReadyConfiguration;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyWidevineConfiguration;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyCompareOperation;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyCondition;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyType;
import com.azure.resourcemanager.mediaservices.models.TrackSelection;
import java.util.Arrays;

/**
 * Samples for StreamingPolicies Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-
     * policies-create-secure-streaming.json
     */
    /**
     * Sample code: Creates a Streaming Policy with secure streaming.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithSecureStreaming(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingPolicies().define("UserCreatedSecureStreamingPolicy")
            .withExistingMediaService("contosorg", "contosomedia")
            .withDefaultContentKeyPolicyName(
                "PolicyWithMultipleOptions")
            .withEnvelopeEncryption(
                new EnvelopeEncryption()
                    .withEnabledProtocols(new EnabledProtocols().withDownload(false).withDash(true).withHls(true)
                        .withSmoothStreaming(true))
                    .withContentKeys(
                        new StreamingPolicyContentKeys().withDefaultKey(new DefaultKey().withLabel("aesDefaultKey")))
                    .withCustomKeyAcquisitionUrlTemplate("fakeTokenPlaceholder"))
            .withCommonEncryptionCenc(new CommonEncryptionCenc()
                .withEnabledProtocols(
                    new EnabledProtocols().withDownload(false).withDash(true).withHls(false).withSmoothStreaming(true))
                .withClearTracks(Arrays.asList(new TrackSelection().withTrackSelections(
                    Arrays.asList(new TrackPropertyCondition().withProperty(TrackPropertyType.FOUR_CC)
                        .withOperation(TrackPropertyCompareOperation.EQUAL).withValue("hev1")))))
                .withContentKeys(new StreamingPolicyContentKeys()
                    .withDefaultKey(new DefaultKey().withLabel("cencDefaultKey")))
                .withDrm(new CencDrmConfiguration().withPlayReady(new StreamingPolicyPlayReadyConfiguration()
                    .withCustomLicenseAcquisitionUrlTemplate(
                        "https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}")
                    .withPlayReadyCustomAttributes("PlayReady CustomAttributes"))
                    .withWidevine(new StreamingPolicyWidevineConfiguration().withCustomLicenseAcquisitionUrlTemplate(
                        "https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId"))))
            .withCommonEncryptionCbcs(new CommonEncryptionCbcs()
                .withEnabledProtocols(
                    new EnabledProtocols().withDownload(false).withDash(false).withHls(true).withSmoothStreaming(false))
                .withContentKeys(
                    new StreamingPolicyContentKeys().withDefaultKey(new DefaultKey().withLabel("cbcsDefaultKey")))
                .withDrm(new CbcsDrmConfiguration().withFairPlay(new StreamingPolicyFairPlayConfiguration()
                    .withCustomLicenseAcquisitionUrlTemplate(
                        "https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}")
                    .withAllowPersistentLicense(true))))
            .create();
    }
}
