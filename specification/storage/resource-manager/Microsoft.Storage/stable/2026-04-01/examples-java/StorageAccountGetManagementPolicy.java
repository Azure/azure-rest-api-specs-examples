
import com.azure.resourcemanager.storage.models.ManagementPolicyName;

/**
 * Samples for ManagementPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountGetManagementPolicy.json
     */
    /**
     * Sample code: StorageAccountGetManagementPolicies.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountGetManagementPolicies(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getManagementPolicies().getWithResponse("res6977", "sto2527",
            ManagementPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
