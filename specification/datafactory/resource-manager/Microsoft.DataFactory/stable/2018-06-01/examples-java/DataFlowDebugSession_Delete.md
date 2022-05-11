Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.15/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.datafactory.models.DeleteDataFlowDebugSessionRequest;

/** Samples for DataFlowDebugSession Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Delete.json
     */
    /**
     * Sample code: DataFlowDebugSession_Delete.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowDebugSessionDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .dataFlowDebugSessions()
            .deleteWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                new DeleteDataFlowDebugSessionRequest().withSessionId("91fb57e0-8292-47be-89ff-c8f2d2bb2a7e"),
                Context.NONE);
    }
}
```
