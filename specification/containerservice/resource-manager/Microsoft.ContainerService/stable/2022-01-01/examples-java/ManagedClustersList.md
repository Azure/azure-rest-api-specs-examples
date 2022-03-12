Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ManagedClusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-01-01/examples/ManagedClustersList.json
     */
    /**
     * Sample code: List Managed Clusters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedClusters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().list(Context.NONE);
    }
}
```
