Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyDefinitions ListByManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicyDefinitionsByManagementGroup.json
     */
    /**
     * Sample code: List policy definitions by management group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPolicyDefinitionsByManagementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyDefinitions()
            .listByManagementGroup("MyManagementGroup", null, null, Context.NONE);
    }
}
```
