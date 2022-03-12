Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for MaintenanceConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-01-01/examples/MaintenanceConfigurationsGet.json
     */
    /**
     * Sample code: Get Maintenance Configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMaintenanceConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getMaintenanceConfigurations()
            .getWithResponse("rg1", "clustername1", "default", Context.NONE);
    }
}
```
