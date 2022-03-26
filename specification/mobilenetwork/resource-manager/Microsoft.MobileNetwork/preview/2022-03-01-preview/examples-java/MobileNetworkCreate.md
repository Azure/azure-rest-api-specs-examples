Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.2/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mobilenetwork.models.PlmnId;

/** Samples for MobileNetworks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/MobileNetworkCreate.json
     */
    /**
     * Sample code: Create mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createMobileNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .mobileNetworks()
            .define("testMobileNetwork")
            .withRegion("eastus")
            .withExistingResourceGroup("rg1")
            .withPublicLandMobileNetworkIdentifier(new PlmnId().withMcc("001").withMnc("01"))
            .create();
    }
}
```
