
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyClearKeyConfiguration;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOption;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyRestrictionTokenType;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicySymmetricTokenKey;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyTokenRestriction;
import java.util.Arrays;

/**
 * Samples for ContentKeyPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-
     * policies-create-nodrm-token.json
     */
    /**
     * Sample code: Creates a Content Key Policy with ClearKey option and Token Restriction.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAContentKeyPolicyWithClearKeyOptionAndTokenRestriction(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.contentKeyPolicies().define("PolicyWithClearKeyOptionAndSwtTokenRestriction")
            .withExistingMediaService("contosorg", "contosomedia").withDescription("ArmPolicyDescription")
            .withOptions(
                Arrays.asList(new ContentKeyPolicyOption().withName("ClearKeyOption")
                    .withConfiguration(new ContentKeyPolicyClearKeyConfiguration())
                    .withRestriction(new ContentKeyPolicyTokenRestriction().withIssuer("urn:issuer")
                        .withAudience("urn:audience")
                        .withPrimaryVerificationKey(
                            new ContentKeyPolicySymmetricTokenKey().withKeyValue("AAAAAAAAAAAAAAAAAAAAAA==".getBytes()))
                        .withRestrictionTokenType(ContentKeyPolicyRestrictionTokenType.SWT))))
            .create();
    }
}
