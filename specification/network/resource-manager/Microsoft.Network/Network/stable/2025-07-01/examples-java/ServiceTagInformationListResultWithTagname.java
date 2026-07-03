
/**
 * Samples for ServiceTagInformation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceTagInformationListResultWithTagname.json
     */
    /**
     * Sample code: Get list of service tags with tag name.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getListOfServiceTagsWithTagName(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceTagInformations().list("westeurope", null, "ApiManagement",
            com.azure.core.util.Context.NONE);
    }
}
