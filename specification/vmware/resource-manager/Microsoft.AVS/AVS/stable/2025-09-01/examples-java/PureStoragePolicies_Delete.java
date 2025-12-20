
/**
 * Samples for PureStoragePolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PureStoragePolicies_Delete.json
     */
    /**
     * Sample code: PureStoragePolicies_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void pureStoragePoliciesDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.pureStoragePolicies().delete("group1", "cloud1", "storagePolicy1", com.azure.core.util.Context.NONE);
    }
}
