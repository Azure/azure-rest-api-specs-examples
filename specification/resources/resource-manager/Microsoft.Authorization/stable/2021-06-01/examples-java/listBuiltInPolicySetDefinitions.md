```java
import com.azure.core.util.Context;

/** Samples for PolicySetDefinitions ListBuiltIn. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listBuiltInPolicySetDefinitions.json
     */
    /**
     * Sample code: List built-in policy set definitions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listBuiltInPolicySetDefinitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicySetDefinitions()
            .listBuiltIn(null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
