import com.azure.resourcemanager.storage.models.ManagementPolicyName;

/** Samples for ManagementPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountDeleteManagementPolicy.json
     */
    /**
     * Sample code: StorageAccountDeleteManagementPolicies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountDeleteManagementPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getManagementPolicies()
            .deleteWithResponse("res6977", "sto2527", ManagementPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
