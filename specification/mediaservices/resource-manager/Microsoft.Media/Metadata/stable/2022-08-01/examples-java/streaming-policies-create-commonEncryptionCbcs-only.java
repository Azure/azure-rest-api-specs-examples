
import com.azure.resourcemanager.mediaservices.models.CbcsDrmConfiguration;
import com.azure.resourcemanager.mediaservices.models.CommonEncryptionCbcs;
import com.azure.resourcemanager.mediaservices.models.DefaultKey;
import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyContentKeys;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyFairPlayConfiguration;

/**
 * Samples for StreamingPolicies Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-
     * policies-create-commonEncryptionCbcs-only.json
     */
    /**
     * Sample code: Creates a Streaming Policy with commonEncryptionCbcs only.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithCommonEncryptionCbcsOnly(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingPolicies().define("UserCreatedSecureStreamingPolicyWithCommonEncryptionCbcsOnly")
            .withExistingMediaService("contosorg", "contosomedia")
            .withDefaultContentKeyPolicyName("PolicyWithMultipleOptions")
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
