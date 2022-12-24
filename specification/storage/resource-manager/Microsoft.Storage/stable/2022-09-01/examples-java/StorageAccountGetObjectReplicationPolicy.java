import com.azure.core.util.Context;

/** Samples for ObjectReplicationPoliciesOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountGetObjectReplicationPolicy.json
     */
    /**
     * Sample code: StorageAccountGetObjectReplicationPolicies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountGetObjectReplicationPolicies(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getObjectReplicationPoliciesOperations()
            .getWithResponse("res6977", "sto2527", "{objectReplicationPolicy-Id}", Context.NONE);
    }
}
