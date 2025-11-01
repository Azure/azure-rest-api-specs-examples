
/**
 * Samples for FileServices GetServiceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/FileServicesGetUsage.json
     */
    /**
     * Sample code: GetFileServiceUsage.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFileServiceUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileServices().getServiceUsageWithResponse("res4410",
            "sto8607", com.azure.core.util.Context.NONE);
    }
}
