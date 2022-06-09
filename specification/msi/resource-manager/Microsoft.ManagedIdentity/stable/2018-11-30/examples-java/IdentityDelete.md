```java
import com.azure.core.util.Context;

/** Samples for UserAssignedIdentities Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2018-11-30/examples/IdentityDelete.json
     */
    /**
     * Sample code: IdentityDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void identityDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .identities()
            .manager()
            .serviceClient()
            .getUserAssignedIdentities()
            .deleteWithResponse("rgName", "resourceName", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
