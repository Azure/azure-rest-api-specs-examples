Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyClearKeyConfiguration;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOption;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyRestrictionTokenType;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicySymmetricTokenKey;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyTokenRestriction;
import java.util.Arrays;

/** Samples for ContentKeyPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-create-nodrm-token.json
     */
    /**
     * Sample code: Creates a Content Key Policy with ClearKey option and Token Restriction.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAContentKeyPolicyWithClearKeyOptionAndTokenRestriction(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .define("PolicyWithClearKeyOptionAndSwtTokenRestriction")
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
                                    .withRestrictionTokenType(ContentKeyPolicyRestrictionTokenType.SWT))))
            .create();
    }
}
```
