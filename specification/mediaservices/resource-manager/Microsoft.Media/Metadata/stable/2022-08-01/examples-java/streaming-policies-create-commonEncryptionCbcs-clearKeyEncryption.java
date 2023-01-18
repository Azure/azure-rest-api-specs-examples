import com.azure.resourcemanager.mediaservices.models.ClearKeyEncryptionConfiguration;
import com.azure.resourcemanager.mediaservices.models.CommonEncryptionCbcs;
import com.azure.resourcemanager.mediaservices.models.DefaultKey;
import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyContentKeys;

/** Samples for StreamingPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-commonEncryptionCbcs-clearKeyEncryption.json
     */
    /**
     * Sample code: Creates a Streaming Policy with ClearKey encryption in commonEncryptionCbcs.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithClearKeyEncryptionInCommonEncryptionCbcs(
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
                    .withClearKeyEncryptionConfiguration(
                        new ClearKeyEncryptionConfiguration()
                            .withCustomKeysAcquisitionUrlTemplate("fakeTokenPlaceholder")))
            .create();
    }
}
