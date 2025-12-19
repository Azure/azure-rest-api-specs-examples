
/**
 * Samples for PureStoragePolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PureStoragePolicies_List.json
     */
    /**
     * Sample code: PureStoragePolicies_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void pureStoragePoliciesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.pureStoragePolicies().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
