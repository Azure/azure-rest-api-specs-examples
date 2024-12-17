
import com.azure.resourcemanager.playwrighttesting.models.QuotaNames;

/**
 * Samples for Quotas Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/Quotas_Get.json
     */
    /**
     * Sample code: Quotas_Get.
     * 
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void quotasGet(com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager.quotas().getWithResponse("eastus", QuotaNames.SCALABLE_EXECUTION, com.azure.core.util.Context.NONE);
    }
}
