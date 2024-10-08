
/**
 * Samples for DatabaseExtensionsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetDatabaseExtensions.json
     */
    /**
     * Sample code: Get database extensions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDatabaseExtensions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseExtensionsOperations().getWithResponse(
            "rg_a1f9d6f8-30d5-4228-9504-8a364361bca3", "srv_65858e0f-b1d1-4bdc-8351-a7da86ca4939",
            "11aa6c5e-58ed-4693-b303-3b8e3131deaa", "polybaseimport", com.azure.core.util.Context.NONE);
    }
}
