Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicySetDefinitions GetAtManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicySetDefinitionAtManagementGroup.json
     */
    /**
     * Sample code: Retrieve a policy set definition at management group level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAPolicySetDefinitionAtManagementGroupLevel(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicySetDefinitions()
            .getAtManagementGroupWithResponse("CostManagement", "MyManagementGroup", Context.NONE);
    }
}
```
