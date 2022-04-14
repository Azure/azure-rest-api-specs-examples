Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-azurestackhci_1.0.0-beta.2/sdk/azurestackhci/azure-resourcemanager-azurestackhci/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Clusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-01-01/examples/DeleteCluster.json
     */
    /**
     * Sample code: Delete cluster.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteCluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().deleteWithResponse("test-rg", "myCluster", Context.NONE);
    }
}
```
