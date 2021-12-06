Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.avs.models.AddonSrmProperties;

/** Samples for Addons CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Addons_CreateOrUpdate_SRM.json
     */
    /**
     * Sample code: Addons_CreateOrUpdate_SRM.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void addonsCreateOrUpdateSRM(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .addons()
            .define("srm")
            .withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new AddonSrmProperties().withLicenseKey("41915178-A8FF-4A4D-B683-6D735AF5E3F5"))
            .create();
    }
}
```
