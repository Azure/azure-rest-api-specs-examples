
/**
 * Samples for ServiceTagInformation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ServiceTagInformationListResultWithNoAddressPrefixes.json
     */
    /**
     * Sample code: Get list of service tags with no address prefixes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getListOfServiceTagsWithNoAddressPrefixes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceTagInformations().list("westeurope", true, null,
            com.azure.core.util.Context.NONE);
    }
}
