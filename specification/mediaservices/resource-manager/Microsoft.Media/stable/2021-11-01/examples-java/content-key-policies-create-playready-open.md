Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOpenRestriction;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOption;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyPlayReadyConfiguration;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyPlayReadyContentType;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyPlayReadyLicense;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyPlayReadyLicenseType;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyPlayReadyPlayRight;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyPlayReadyUnknownOutputPassingOption;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for ContentKeyPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-create-playready-open.json
     */
    /**
     * Sample code: Creates a Content Key Policy with PlayReady option and Open Restriction.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAContentKeyPolicyWithPlayReadyOptionAndOpenRestriction(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .define("PolicyWithPlayReadyOptionAndOpenRestriction")
            .withExistingMediaService("contoso", "contosomedia")
            .withDescription("ArmPolicyDescription")
            .withOptions(
                Arrays
                    .asList(
                        new ContentKeyPolicyOption()
                            .withName("ArmPolicyOptionName")
                            .withConfiguration(
                                new ContentKeyPolicyPlayReadyConfiguration()
                                    .withLicenses(
                                        Arrays
                                            .asList(
                                                new ContentKeyPolicyPlayReadyLicense()
                                                    .withAllowTestDevices(true)
                                                    .withBeginDate(OffsetDateTime.parse("2017-10-16T18:22:53.46Z"))
                                                    .withPlayRight(
                                                        new ContentKeyPolicyPlayReadyPlayRight()
                                                            .withScmsRestriction(2)
                                                            .withDigitalVideoOnlyContentRestriction(false)
                                                            .withImageConstraintForAnalogComponentVideoRestriction(true)
                                                            .withImageConstraintForAnalogComputerMonitorRestriction(
                                                                false)
                                                            .withAllowPassingVideoContentToUnknownOutput(
                                                                ContentKeyPolicyPlayReadyUnknownOutputPassingOption
                                                                    .NOT_ALLOWED))
                                                    .withLicenseType(ContentKeyPolicyPlayReadyLicenseType.PERSISTENT)
                                                    .withContentKeyLocation(
                                                        new ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader())
                                                    .withContentType(
                                                        ContentKeyPolicyPlayReadyContentType.ULTRA_VIOLET_DOWNLOAD))))
                            .withRestriction(new ContentKeyPolicyOpenRestriction())))
            .create();
    }
}
```
