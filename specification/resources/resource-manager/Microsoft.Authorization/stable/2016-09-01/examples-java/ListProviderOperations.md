Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AuthorizationOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2016-09-01/examples/ListProviderOperations.json
     */
    /**
     * Sample code: List provider operations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProviderOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getAuthorizationOperations().list(Context.NONE);
    }
}
```
