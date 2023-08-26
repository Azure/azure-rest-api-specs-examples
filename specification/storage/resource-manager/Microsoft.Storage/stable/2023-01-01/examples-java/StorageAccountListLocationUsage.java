/** Samples for Usages ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountListLocationUsage.json
     */
    /**
     * Sample code: UsageList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void usageList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getUsages()
            .listByLocation("eastus2(stage)", com.azure.core.util.Context.NONE);
    }
}
