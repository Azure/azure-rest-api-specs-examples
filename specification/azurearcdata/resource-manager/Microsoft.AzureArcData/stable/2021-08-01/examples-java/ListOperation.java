
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/ListOperation.json
     */
    /**
     * Sample code: Lists all of the available Azure Data Services on Azure Arc API operations.
     * 
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void listsAllOfTheAvailableAzureDataServicesOnAzureArcAPIOperations(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
