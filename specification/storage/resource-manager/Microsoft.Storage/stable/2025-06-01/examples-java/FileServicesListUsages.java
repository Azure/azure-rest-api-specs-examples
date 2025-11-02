
/**
 * Samples for FileServices ListServiceUsages.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/FileServicesListUsages.json
     */
    /**
     * Sample code: ListFileServiceUsages.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listFileServiceUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileServices().listServiceUsages("res4410", "sto8607",
            null, com.azure.core.util.Context.NONE);
    }
}
