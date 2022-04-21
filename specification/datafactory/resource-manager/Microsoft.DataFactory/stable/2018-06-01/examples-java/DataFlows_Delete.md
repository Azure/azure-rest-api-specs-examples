Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.14/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DataFlows Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Delete.json
     */
    /**
     * Sample code: DataFlows_Delete.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowsDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .dataFlows()
            .deleteWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleDataFlow", Context.NONE);
    }
}
```
