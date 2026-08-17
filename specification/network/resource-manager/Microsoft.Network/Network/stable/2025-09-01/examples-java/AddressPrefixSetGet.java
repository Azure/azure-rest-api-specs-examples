
/**
 * Samples for AddressPrefixSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/AddressPrefixSetGet.json
     */
    /**
     * Sample code: Get address prefix set.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAddressPrefixSet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAddressPrefixSets().getWithResponse("rg1", "test-asg", "test-prefix-set",
            com.azure.core.util.Context.NONE);
    }
}
