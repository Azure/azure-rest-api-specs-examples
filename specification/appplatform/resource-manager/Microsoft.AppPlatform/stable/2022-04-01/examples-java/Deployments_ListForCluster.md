Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Deployments ListForCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_ListForCluster.json
     */
    /**
     * Sample code: Deployments_ListForCluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsListForCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .listForCluster("myResourceGroup", "myservice", null, Context.NONE);
    }
}
```
