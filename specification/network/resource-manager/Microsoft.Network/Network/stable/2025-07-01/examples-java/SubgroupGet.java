
/**
 * Samples for Subgroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubgroupGet.json
     */
    /**
     * Sample code: Get subgroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSubgroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubgroups().getWithResponse("rg1", "test-ig", "subgroup0",
            com.azure.core.util.Context.NONE);
    }
}
