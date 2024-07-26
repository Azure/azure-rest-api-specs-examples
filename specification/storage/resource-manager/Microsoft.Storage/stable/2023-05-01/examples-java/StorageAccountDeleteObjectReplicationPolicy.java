
/**
 * Samples for ObjectReplicationPoliciesOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * StorageAccountDeleteObjectReplicationPolicy.json
     */
    /**
     * Sample code: StorageAccountDeleteObjectReplicationPolicies.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        storageAccountDeleteObjectReplicationPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getObjectReplicationPoliciesOperations()
            .deleteWithResponse("res6977", "sto2527", "{objectReplicationPolicy-Id}", com.azure.core.util.Context.NONE);
    }
}
