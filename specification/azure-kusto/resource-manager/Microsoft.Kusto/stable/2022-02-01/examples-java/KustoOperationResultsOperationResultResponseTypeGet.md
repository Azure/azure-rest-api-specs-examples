```java
import com.azure.core.util.Context;

/** Samples for OperationsResultsLocation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoOperationResultsOperationResultResponseTypeGet.json
     */
    /**
     * Sample code: KustoOperationsResultsLocationGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoOperationsResultsLocationGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .operationsResultsLocations()
            .getWithResponse("westus", "30972f1b-b61d-4fd8-bd34-3dcfa24670f3", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
