
import com.azure.resourcemanager.storage.models.ManagementPolicyName;

/**
 * Samples for ManagementPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountDeleteManagementPolicy.json
     */
    /**
     * Sample code: StorageAccountDeleteManagementPolicies.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountDeleteManagementPolicies(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getManagementPolicies().deleteWithResponse("res6977", "sto2527",
            ManagementPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
