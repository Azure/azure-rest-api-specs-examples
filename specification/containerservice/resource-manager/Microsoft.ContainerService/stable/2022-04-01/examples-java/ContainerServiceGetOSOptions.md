Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ManagedClusters GetOSOptions. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-04-01/examples/ContainerServiceGetOSOptions.json
     */
    /**
     * Sample code: Get Container Service OS Options.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getContainerServiceOSOptions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .getOSOptionsWithResponse("location1", null, Context.NONE);
    }
}
```
