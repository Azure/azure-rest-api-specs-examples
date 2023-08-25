/** Samples for ObjectReplicationPoliciesOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountListObjectReplicationPolicies.json
     */
    /**
     * Sample code: StorageAccountListObjectReplicationPolicies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListObjectReplicationPolicies(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getObjectReplicationPoliciesOperations()
            .list("res6977", "sto2527", com.azure.core.util.Context.NONE);
    }
}
