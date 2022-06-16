import com.azure.core.util.Context;

/** Samples for TableServices GetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/TableServicesGet.json
     */
    /**
     * Sample code: TableServicesGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tableServicesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getTableServices()
            .getServicePropertiesWithResponse("res4410", "sto8607", Context.NONE);
    }
}
