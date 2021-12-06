Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.8/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Datasets ListByFactory. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_ListByFactory.json
     */
    /**
     * Sample code: Datasets_ListByFactory.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void datasetsListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.datasets().listByFactory("exampleResourceGroup", "exampleFactoryName", Context.NONE);
    }
}
```
