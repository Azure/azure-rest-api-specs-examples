Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyExemptions ListForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForResource.json
     */
    /**
     * Sample code: List all policy exemptions that apply to a resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllPolicyExemptionsThatApplyToAResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyExemptions()
            .listForResource(
                "TestResourceGroup",
                "Microsoft.Compute",
                "virtualMachines/MyTestVm",
                "domainNames",
                "MyTestComputer.cloudapp.net",
                null,
                Context.NONE);
    }
}
```
