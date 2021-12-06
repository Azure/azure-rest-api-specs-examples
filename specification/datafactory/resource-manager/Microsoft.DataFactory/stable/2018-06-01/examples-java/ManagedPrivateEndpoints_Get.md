Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.8/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ManagedPrivateEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Get.json
     */
    /**
     * Sample code: ManagedPrivateEndpoints_Get.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void managedPrivateEndpointsGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .managedPrivateEndpoints()
            .getWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                "exampleManagedVirtualNetworkName",
                "exampleManagedPrivateEndpointName",
                null,
                Context.NONE);
    }
}
```
