```java
import com.azure.core.util.Context;

/** Samples for Services Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/ServiceGet.json
     */
    /**
     * Sample code: Get service.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getService(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.services().getWithResponse("rg1", "testMobileNetwork", "TestService", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.2/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.
