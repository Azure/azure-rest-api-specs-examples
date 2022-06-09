```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.edgeorder.models.CancellationReason;

/** Samples for ResourceProvider CancelOrderItem. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CancelOrderItem.json
     */
    /**
     * Sample code: CancelOrderItem.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void cancelOrderItem(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager
            .resourceProviders()
            .cancelOrderItemWithResponse(
                "TestOrderItemName1", "TestRG", new CancellationReason().withReason("Order cancelled"), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-edgeorder_1.0.0-beta.1/sdk/edgeorder/azure-resourcemanager-edgeorder/README.md) on how to add the SDK to your project and authenticate.
