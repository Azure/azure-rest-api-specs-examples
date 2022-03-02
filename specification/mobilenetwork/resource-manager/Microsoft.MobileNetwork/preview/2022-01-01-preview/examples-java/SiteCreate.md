Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.1/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import java.util.Arrays;

/** Samples for Sites CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SiteCreate.json
     */
    /**
     * Sample code: Create mobile network site.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createMobileNetworkSite(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .sites()
            .define("testSite")
            .withRegion("testLocation")
            .withExistingMobileNetwork("rg1", "testMobileNetwork")
            .withNetworkFunctions(
                Arrays
                    .asList(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HybridNetwork/networkFunctions/testNf")))
            .create();
    }
}
```
