
import com.azure.resourcemanager.playwrighttesting.models.QuotaNames;

/**
 * Samples for AccountQuotas Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/AccountQuotas_Get.json
     */
    /**
     * Sample code: AccountQuotas_Get.
     * 
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void accountQuotasGet(com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager.accountQuotas().getWithResponse("dummyrg", "myPlaywrightAccount", QuotaNames.SCALABLE_EXECUTION,
            com.azure.core.util.Context.NONE);
    }
}
