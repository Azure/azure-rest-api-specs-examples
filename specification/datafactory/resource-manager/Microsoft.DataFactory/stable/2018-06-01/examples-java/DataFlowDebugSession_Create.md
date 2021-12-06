Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.8/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.datafactory.models.CreateDataFlowDebugSessionRequest;
import com.azure.resourcemanager.datafactory.models.DataFlowComputeType;
import com.azure.resourcemanager.datafactory.models.IntegrationRuntimeComputeProperties;
import com.azure.resourcemanager.datafactory.models.IntegrationRuntimeDataFlowProperties;
import com.azure.resourcemanager.datafactory.models.IntegrationRuntimeDebugResource;
import com.azure.resourcemanager.datafactory.models.ManagedIntegrationRuntime;
import java.util.HashMap;
import java.util.Map;

/** Samples for DataFlowDebugSession Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Create.json
     */
    /**
     * Sample code: DataFlowDebugSession_Create.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowDebugSessionCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .dataFlowDebugSessions()
            .create(
                "exampleResourceGroup",
                "exampleFactoryName",
                new CreateDataFlowDebugSessionRequest()
                    .withTimeToLive(60)
                    .withIntegrationRuntime(
                        new IntegrationRuntimeDebugResource()
                            .withName("ir1")
                            .withProperties(
                                new ManagedIntegrationRuntime()
                                    .withComputeProperties(
                                        new IntegrationRuntimeComputeProperties()
                                            .withLocation("AutoResolve")
                                            .withDataFlowProperties(
                                                new IntegrationRuntimeDataFlowProperties()
                                                    .withComputeType(DataFlowComputeType.GENERAL)
                                                    .withCoreCount(48)
                                                    .withTimeToLive(10)
                                                    .withAdditionalProperties(mapOf()))
                                            .withAdditionalProperties(mapOf())))),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
