Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyDefinitions GetBuiltIn. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getBuiltinPolicyDefinition.json
     */
    /**
     * Sample code: Retrieve a built-in policy definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveABuiltInPolicyDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyDefinitions()
            .getBuiltInWithResponse("7433c107-6db4-4ad1-b57a-a76dce0154a1", Context.NONE);
    }
}
```
