Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyExemptions ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForResourceGroup.json
     */
    /**
     * Sample code: List policy exemptions that apply to a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPolicyExemptionsThatApplyToAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyExemptions()
            .listByResourceGroup("TestResourceGroup", "atScope()", Context.NONE);
    }
}
```
