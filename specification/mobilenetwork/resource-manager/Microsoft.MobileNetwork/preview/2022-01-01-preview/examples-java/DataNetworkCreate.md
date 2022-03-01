Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.1/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for DataNetworks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/DataNetworkCreate.json
     */
    /**
     * Sample code: Create mobile network dataNetwork.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createMobileNetworkDataNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .dataNetworks()
            .define("testDataNetwork")
            .withRegion("eastus")
            .withExistingMobileNetwork("rg1", "testMobileNetwork")
            .withDescription("myFavouriteDataNetwork")
            .create();
    }
}
```
