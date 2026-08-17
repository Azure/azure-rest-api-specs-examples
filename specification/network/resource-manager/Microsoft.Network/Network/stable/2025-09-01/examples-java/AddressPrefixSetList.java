
/**
 * Samples for AddressPrefixSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/AddressPrefixSetList.json
     */
    /**
     * Sample code: List address prefix sets.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAddressPrefixSets(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAddressPrefixSets().list("rg1", "test-asg", com.azure.core.util.Context.NONE);
    }
}
