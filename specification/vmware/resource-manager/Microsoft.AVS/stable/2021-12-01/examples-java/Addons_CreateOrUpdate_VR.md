```java
import com.azure.resourcemanager.avs.models.AddonVrProperties;

/** Samples for Addons CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Addons_CreateOrUpdate_VR.json
     */
    /**
     * Sample code: Addons_CreateOrUpdate_VR.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void addonsCreateOrUpdateVR(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .addons()
            .define("vr")
            .withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new AddonVrProperties().withVrsCount(1))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.
