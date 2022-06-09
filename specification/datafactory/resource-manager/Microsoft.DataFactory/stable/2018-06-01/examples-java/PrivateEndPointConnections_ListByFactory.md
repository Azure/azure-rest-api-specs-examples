```java
import com.azure.core.util.Context;

/** Samples for PrivateEndPointConnections ListByFactory. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PrivateEndPointConnections_ListByFactory.json
     */
    /**
     * Sample code: privateEndPointConnections_ListByFactory.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void privateEndPointConnectionsListByFactory(
        com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.privateEndPointConnections().listByFactory("exampleResourceGroup", "exampleFactoryName", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.15/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.
