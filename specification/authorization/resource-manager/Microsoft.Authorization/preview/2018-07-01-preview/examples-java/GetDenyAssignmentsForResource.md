Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DenyAssignments ListForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-07-01-preview/examples/GetDenyAssignmentsForResource.json
     */
    /**
     * Sample code: List deny assignments for resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDenyAssignmentsForResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getDenyAssignments()
            .listForResource(
                "rgname",
                "resourceProviderNamespace",
                "parentResourcePath",
                "resourceType",
                "resourceName",
                null,
                Context.NONE);
    }
}
```
