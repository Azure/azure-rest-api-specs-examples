/** Samples for BlobServices GetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobServicesGet.json
     */
    /**
     * Sample code: GetBlobServices.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getBlobServices(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobServices()
            .getServicePropertiesWithResponse("res4410", "sto8607", com.azure.core.util.Context.NONE);
    }
}
