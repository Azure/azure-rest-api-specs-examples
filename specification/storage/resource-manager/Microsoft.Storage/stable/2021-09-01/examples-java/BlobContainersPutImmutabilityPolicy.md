Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.ImmutabilityPolicyInner;

/** Samples for BlobContainers CreateOrUpdateImmutabilityPolicy. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/BlobContainersPutImmutabilityPolicy.json
     */
    /**
     * Sample code: CreateOrUpdateImmutabilityPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateImmutabilityPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .createOrUpdateImmutabilityPolicyWithResponse(
                "res1782",
                "sto7069",
                "container6397",
                null,
                new ImmutabilityPolicyInner()
                    .withImmutabilityPeriodSinceCreationInDays(3)
                    .withAllowProtectedAppendWrites(true),
                Context.NONE);
    }
}
```
