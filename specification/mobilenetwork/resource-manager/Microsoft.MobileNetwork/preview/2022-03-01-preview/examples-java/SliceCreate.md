Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.2/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mobilenetwork.models.Snssai;

/** Samples for Slices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SliceCreate.json
     */
    /**
     * Sample code: Create mobile network slice.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createMobileNetworkSlice(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .slices()
            .define("testSlice")
            .withRegion("eastus")
            .withExistingMobileNetwork("rg1", "testMobileNetwork")
            .withSnssai(new Snssai().withSst(1).withSd("1abcde"))
            .withDescription("myFavouriteSlice")
            .create();
    }
}
```
