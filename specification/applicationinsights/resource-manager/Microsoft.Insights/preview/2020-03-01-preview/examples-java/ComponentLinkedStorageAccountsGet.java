import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.StorageType;

/** Samples for ComponentLinkedStorageAccountsOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsGet.json
     */
    /**
     * Sample code: ComponentLinkedStorageAccountsGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentLinkedStorageAccountsGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .componentLinkedStorageAccountsOperations()
            .getWithResponse("someResourceGroupName", "myComponent", StorageType.SERVICE_PROFILER, Context.NONE);
    }
}
