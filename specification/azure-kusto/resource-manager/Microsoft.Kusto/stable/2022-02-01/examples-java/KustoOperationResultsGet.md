Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for OperationsResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoOperationResultsGet.json
     */
    /**
     * Sample code: KustoOperationResultsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoOperationResultsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.operationsResults().getWithResponse("westus", "30972f1b-b61d-4fd8-bd34-3dcfa24670f3", Context.NONE);
    }
}
```
