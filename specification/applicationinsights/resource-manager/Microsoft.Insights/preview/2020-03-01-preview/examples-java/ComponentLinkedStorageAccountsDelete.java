
import com.azure.resourcemanager.applicationinsights.models.StorageType;

/**
 * Samples for ComponentLinkedStorageAccountsOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/
     * ComponentLinkedStorageAccountsDelete.json
     */
    /**
     * Sample code: ComponentLinkedStorageAccountsDelete.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentLinkedStorageAccountsDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.componentLinkedStorageAccountsOperations().deleteWithResponse("someResourceGroupName", "myComponent",
            StorageType.SERVICE_PROFILER, com.azure.core.util.Context.NONE);
    }
}
