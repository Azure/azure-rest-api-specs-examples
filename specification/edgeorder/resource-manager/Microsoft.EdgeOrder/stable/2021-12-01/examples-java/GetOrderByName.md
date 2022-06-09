```java
import com.azure.core.util.Context;

/** Samples for ResourceProvider GetOrderByName. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetOrderByName.json
     */
    /**
     * Sample code: GetOrderByName.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void getOrderByName(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager
            .resourceProviders()
            .getOrderByNameWithResponse("TestOrderItemName901", "TestRG", "%7B%7B%7Blocation%7D%7D", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-edgeorder_1.0.0-beta.1/sdk/edgeorder/azure-resourcemanager-edgeorder/README.md) on how to add the SDK to your project and authenticate.
