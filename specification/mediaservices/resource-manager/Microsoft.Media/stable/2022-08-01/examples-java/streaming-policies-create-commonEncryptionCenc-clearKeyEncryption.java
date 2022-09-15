import com.azure.resourcemanager.mediaservices.models.ClearKeyEncryptionConfiguration;
import com.azure.resourcemanager.mediaservices.models.CommonEncryptionCenc;
import com.azure.resourcemanager.mediaservices.models.DefaultKey;
import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.StreamingPolicyContentKeys;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyCompareOperation;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyCondition;
import com.azure.resourcemanager.mediaservices.models.TrackPropertyType;
import com.azure.resourcemanager.mediaservices.models.TrackSelection;
import java.util.Arrays;

/** Samples for StreamingPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-policies-create-commonEncryptionCenc-clearKeyEncryption.json
     */
    /**
     * Sample code: Creates a Streaming Policy with ClearKey encryption in commonEncryptionCenc.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithClearKeyEncryptionInCommonEncryptionCenc(
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
                    .withClearKeyEncryptionConfiguration(
                        new ClearKeyEncryptionConfiguration()
                            .withCustomKeysAcquisitionUrlTemplate(
                                "https://contoso.com/{AlternativeMediaId}/clearkey/")))
            .create();
    }
}
