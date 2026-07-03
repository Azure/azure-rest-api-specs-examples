
/**
 * Samples for ServiceTags List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceTagsList.json
     */
    /**
     * Sample code: Get list of service tags.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getListOfServiceTags(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceTags().listWithResponse("westcentralus", com.azure.core.util.Context.NONE);
    }
}
