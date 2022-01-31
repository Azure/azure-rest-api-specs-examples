Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.ComponentLinkedStorageAccounts;
import com.azure.resourcemanager.applicationinsights.models.StorageType;

/** Samples for ComponentLinkedStorageAccountsOperation Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsUpdate.json
     */
    /**
     * Sample code: ComponentLinkedStorageAccountsUpdate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentLinkedStorageAccountsUpdate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        ComponentLinkedStorageAccounts resource =
            manager
                .componentLinkedStorageAccountsOperations()
                .getWithResponse("someResourceGroupName", "myComponent", StorageType.SERVICE_PROFILER, Context.NONE)
                .getValue();
        resource
            .update()
            .withLinkedStorageAccount(
                "/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/someResourceGroupName/providers/Microsoft.Storage/storageAccounts/storageaccountname")
            .apply();
    }
}
```
