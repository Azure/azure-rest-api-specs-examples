
import com.azure.resourcemanager.applicationinsights.models.StorageType;

/**
 * Samples for ComponentLinkedStorageAccountsOperation CreateAndUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/
     * ComponentLinkedStorageAccountsCreateAndUpdate.json
     */
    /**
     * Sample code: ComponentLinkedStorageAccountsCreateAndUpdate.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentLinkedStorageAccountsCreateAndUpdate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.componentLinkedStorageAccountsOperations().define(StorageType.SERVICE_PROFILER)
            .withExistingComponent("someResourceGroupName", "myComponent")
            .withLinkedStorageAccount(
                "/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/someResourceGroupName/providers/Microsoft.Storage/storageAccounts/storageaccountname")
            .create();
    }
}
