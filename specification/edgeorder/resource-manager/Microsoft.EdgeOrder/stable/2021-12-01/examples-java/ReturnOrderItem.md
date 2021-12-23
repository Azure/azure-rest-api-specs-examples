Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-edgeorder_1.0.0-beta.1/sdk/edgeorder/azure-resourcemanager-edgeorder/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.edgeorder.models.ReturnOrderItemDetails;

/** Samples for ResourceProvider ReturnOrderItem. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ReturnOrderItem.json
     */
    /**
     * Sample code: ReturnOrderItem.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void returnOrderItem(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager
            .resourceProviders()
            .returnOrderItem(
                "TestOrderName1",
                "TestRG",
                new ReturnOrderItemDetails().withReturnReason("Order returned"),
                Context.NONE);
    }
}
```
