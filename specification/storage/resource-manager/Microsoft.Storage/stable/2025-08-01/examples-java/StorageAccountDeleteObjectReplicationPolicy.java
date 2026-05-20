
/**
 * Samples for ObjectReplicationPoliciesOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountDeleteObjectReplicationPolicy.json
     */
    /**
     * Sample code: StorageAccountDeleteObjectReplicationPolicies.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountDeleteObjectReplicationPolicies(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getObjectReplicationPoliciesOperations().deleteWithResponse("res6977", "sto2527",
            "{objectReplicationPolicy-Id}", com.azure.core.util.Context.NONE);
    }
}
