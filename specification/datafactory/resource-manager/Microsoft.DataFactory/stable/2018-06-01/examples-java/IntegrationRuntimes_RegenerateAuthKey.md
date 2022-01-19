Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.10/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.datafactory.models.IntegrationRuntimeAuthKeyName;
import com.azure.resourcemanager.datafactory.models.IntegrationRuntimeRegenerateKeyParameters;

/** Samples for IntegrationRuntimes RegenerateAuthKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_RegenerateAuthKey.json
     */
    /**
     * Sample code: IntegrationRuntimes_RegenerateAuthKey.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesRegenerateAuthKey(
        com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .integrationRuntimes()
            .regenerateAuthKeyWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                "exampleIntegrationRuntime",
                new IntegrationRuntimeRegenerateKeyParameters().withKeyName(IntegrationRuntimeAuthKeyName.AUTH_KEY2),
                Context.NONE);
    }
}
```
