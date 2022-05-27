Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DataPolicyManifests List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/listDataPolicyManifestsNamespaceFilter.json
     */
    /**
     * Sample code: List data policy manifests with namespace filter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDataPolicyManifestsWithNamespaceFilter(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getDataPolicyManifests()
            .list("namespace eq 'Microsoft.KeyVault'", Context.NONE);
    }
}
```
