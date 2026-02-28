
/**
 * Samples for ServiceTagInformation List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ServiceTagInformationListResult.json
     */
    /**
     * Sample code: Get list of service tags.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getListOfServiceTags(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceTagInformations().list("westeurope", null, null,
            com.azure.core.util.Context.NONE);
    }
}
