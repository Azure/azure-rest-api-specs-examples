import com.azure.core.util.Context;

/** Samples for FileServices GetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileServicesGet.json
     */
    /**
     * Sample code: GetFileServices.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFileServices(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileServices()
            .getServicePropertiesWithResponse("res4410", "sto8607", Context.NONE);
    }
}
