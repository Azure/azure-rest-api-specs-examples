import com.azure.core.util.Context;

/** Samples for ServiceTagInformation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ServiceTagInformationListResultWithNoAddressPrefixes.json
     */
    /**
     * Sample code: Get list of service tags with no address prefixes.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getListOfServiceTagsWithNoAddressPrefixes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getServiceTagInformations()
            .list("westeurope", true, null, Context.NONE);
    }
}
