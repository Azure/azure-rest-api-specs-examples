Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyAssignments ListForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicyAssignmentsForResource.json
     */
    /**
     * Sample code: List all policy assignments that apply to a resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllPolicyAssignmentsThatApplyToAResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyAssignments()
            .listForResource(
                "TestResourceGroup",
                "Microsoft.Compute",
                "virtualMachines/MyTestVm",
                "domainNames",
                "MyTestComputer.cloudapp.net",
                null,
                null,
                Context.NONE);
    }
}
```
