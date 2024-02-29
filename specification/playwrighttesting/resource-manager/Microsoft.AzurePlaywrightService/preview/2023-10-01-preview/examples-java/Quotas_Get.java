
import com.azure.resourcemanager.playwrighttesting.models.QuotaNames;

/**
 * Samples for Quotas Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/
     * examples/Quotas_Get.json
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
