
/**
 * Samples for PureStoragePolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PureStoragePolicies_Get.json
     */
    /**
     * Sample code: PureStoragePolicies_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void pureStoragePoliciesGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.pureStoragePolicies().getWithResponse("group1", "cloud1", "storagePolicy1",
            com.azure.core.util.Context.NONE);
    }
}
