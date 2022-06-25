import com.azure.core.util.Context;

/** Samples for ServiceTagInformation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ServiceTagInformationListResultWithTagname.json
     */
    /**
     * Sample code: Get list of service tags with tag name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getListOfServiceTagsWithTagName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getServiceTagInformations()
            .list("westeurope", null, "ApiManagement", Context.NONE);
    }
}
