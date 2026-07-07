
/**
 * Samples for ObjectReplicationPoliciesOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountListObjectReplicationPolicies.json
     */
    /**
     * Sample code: StorageAccountListObjectReplicationPolicies.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountListObjectReplicationPolicies(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getObjectReplicationPoliciesOperations().list("res6977", "sto2527",
            com.azure.core.util.Context.NONE);
    }
}
