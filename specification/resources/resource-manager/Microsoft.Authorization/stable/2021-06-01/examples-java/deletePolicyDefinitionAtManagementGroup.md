Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyDefinitions DeleteAtManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyDefinitionAtManagementGroup.json
     */
    /**
     * Sample code: Delete a policy definition at management group level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPolicyDefinitionAtManagementGroupLevel(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyDefinitions()
            .deleteAtManagementGroupWithResponse("ResourceNaming", "MyManagementGroup", Context.NONE);
    }
}
```
