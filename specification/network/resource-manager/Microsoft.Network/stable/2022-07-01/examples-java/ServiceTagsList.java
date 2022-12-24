import com.azure.core.util.Context;

/** Samples for ServiceTags List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ServiceTagsList.json
     */
    /**
     * Sample code: Get list of service tags.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getListOfServiceTags(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceTags().listWithResponse("westcentralus", Context.NONE);
    }
}
