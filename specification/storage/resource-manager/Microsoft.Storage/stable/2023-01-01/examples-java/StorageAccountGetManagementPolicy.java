
import com.azure.resourcemanager.storage.models.ManagementPolicyName;

/** Samples for ManagementPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/
     * StorageAccountGetManagementPolicy.json
     */
    /**
     * Sample code: StorageAccountGetManagementPolicies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountGetManagementPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getManagementPolicies().getWithResponse("res6977", "sto2527",
            ManagementPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
