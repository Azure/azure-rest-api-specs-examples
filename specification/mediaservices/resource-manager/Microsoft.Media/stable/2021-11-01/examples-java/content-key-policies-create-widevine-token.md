Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOption;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyRestrictionTokenType;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyRsaTokenKey;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicySymmetricTokenKey;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyTokenRestriction;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyWidevineConfiguration;
import java.util.Arrays;

/** Samples for ContentKeyPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-create-widevine-token.json
     */
    /**
     * Sample code: Creates a Content Key Policy with Widevine option and Token Restriction.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAContentKeyPolicyWithWidevineOptionAndTokenRestriction(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .define("PolicyWithWidevineOptionAndJwtTokenRestriction")
            .withExistingMediaService("contoso", "contosomedia")
            .withDescription("ArmPolicyDescription")
            .withOptions(
                Arrays
                    .asList(
                        new ContentKeyPolicyOption()
                            .withName("widevineoption")
                            .withConfiguration(
                                new ContentKeyPolicyWidevineConfiguration()
                                    .withWidevineTemplate(
                                        "{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"))
                            .withRestriction(
                                new ContentKeyPolicyTokenRestriction()
                                    .withIssuer("urn:issuer")
                                    .withAudience("urn:audience")
                                    .withPrimaryVerificationKey(
                                        new ContentKeyPolicyRsaTokenKey()
                                            .withExponent("AQAB".getBytes())
                                            .withModulus("AQAD".getBytes()))
                                    .withAlternateVerificationKeys(
                                        Arrays
                                            .asList(
                                                new ContentKeyPolicySymmetricTokenKey()
                                                    .withKeyValue("AAAAAAAAAAAAAAAAAAAAAA==".getBytes())))
                                    .withRestrictionTokenType(ContentKeyPolicyRestrictionTokenType.JWT))))
            .create();
    }
}
```
