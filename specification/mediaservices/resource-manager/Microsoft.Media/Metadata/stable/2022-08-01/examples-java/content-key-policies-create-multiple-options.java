import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyClearKeyConfiguration;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOpenRestriction;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOption;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyRestrictionTokenType;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicySymmetricTokenKey;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyTokenRestriction;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyWidevineConfiguration;
import java.util.Arrays;

/** Samples for ContentKeyPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-multiple-options.json
     */
    /**
     * Sample code: Creates a Content Key Policy with multiple options.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAContentKeyPolicyWithMultipleOptions(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .define("PolicyCreatedWithMultipleOptions")
            .withExistingMediaService("contoso", "contosomedia")
            .withDescription("ArmPolicyDescription")
            .withOptions(
                Arrays
                    .asList(
                        new ContentKeyPolicyOption()
                            .withName("ClearKeyOption")
                            .withConfiguration(new ContentKeyPolicyClearKeyConfiguration())
                            .withRestriction(
                                new ContentKeyPolicyTokenRestriction()
                                    .withIssuer("urn:issuer")
                                    .withAudience("urn:audience")
                                    .withPrimaryVerificationKey(
                                        new ContentKeyPolicySymmetricTokenKey()
                                            .withKeyValue("AAAAAAAAAAAAAAAAAAAAAA==".getBytes()))
                                    .withRestrictionTokenType(ContentKeyPolicyRestrictionTokenType.SWT)),
                        new ContentKeyPolicyOption()
                            .withName("widevineoption")
                            .withConfiguration(
                                new ContentKeyPolicyWidevineConfiguration()
                                    .withWidevineTemplate(
                                        "{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"))
                            .withRestriction(new ContentKeyPolicyOpenRestriction())))
            .create();
    }
}
