Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Usages ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/StorageAccountListLocationUsage.json
     */
    /**
     * Sample code: UsageList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void usageList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getUsages().listByLocation("eastus2(stage)", Context.NONE);
    }
}
```
