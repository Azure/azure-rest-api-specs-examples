
/**
 * Samples for ObjectReplicationPoliciesOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountGetObjectReplicationPolicy.json
     */
    /**
     * Sample code: StorageAccountGetObjectReplicationPolicies.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountGetObjectReplicationPolicies(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getObjectReplicationPoliciesOperations().getWithResponse("res6977", "sto2527",
            "{objectReplicationPolicy-Id}", com.azure.core.util.Context.NONE);
    }
}
