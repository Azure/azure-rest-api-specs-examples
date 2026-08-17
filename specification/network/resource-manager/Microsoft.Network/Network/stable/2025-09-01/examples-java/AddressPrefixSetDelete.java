
/**
 * Samples for AddressPrefixSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/AddressPrefixSetDelete.json
     */
    /**
     * Sample code: Delete address prefix set.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteAddressPrefixSet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAddressPrefixSets().delete("rg1", "test-asg", "test-prefix-set",
            com.azure.core.util.Context.NONE);
    }
}
