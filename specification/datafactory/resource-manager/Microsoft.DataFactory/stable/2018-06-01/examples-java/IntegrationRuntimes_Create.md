Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.14/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.datafactory.models.SelfHostedIntegrationRuntime;

/** Samples for IntegrationRuntimes CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Create.json
     */
    /**
     * Sample code: IntegrationRuntimes_Create.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .integrationRuntimes()
            .define("exampleIntegrationRuntime")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withProperties(new SelfHostedIntegrationRuntime().withDescription("A selfhosted integration runtime"))
            .create();
    }
}
```
