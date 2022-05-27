Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Vaults GetDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2019-09-01/examples/getDeletedVault.json
     */
    /**
     * Sample code: Retrieve a deleted vault.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveADeletedVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getVaults()
            .getDeletedWithResponse("sample-vault", "westus", Context.NONE);
    }
}
```
