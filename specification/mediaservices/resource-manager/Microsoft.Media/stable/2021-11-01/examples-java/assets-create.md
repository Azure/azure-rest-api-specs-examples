Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for Assets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-create.json
     */
    /**
     * Sample code: Create an Asset.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createAnAsset(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assets()
            .define("ClimbingMountLogan")
            .withExistingMediaService("contoso", "contosomedia")
            .withDescription("A documentary showing the ascent of Mount Logan")
            .withStorageAccountName("storage0")
            .create();
    }
}
```
