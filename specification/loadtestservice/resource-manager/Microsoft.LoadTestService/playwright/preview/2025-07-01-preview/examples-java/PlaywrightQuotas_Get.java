
import com.azure.resourcemanager.playwright.models.QuotaName;

/**
 * Samples for PlaywrightQuotas Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/PlaywrightQuotas_Get.json
     */
    /**
     * Sample code: PlaywrightQuotas_Get.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void playwrightQuotasGet(com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.playwrightQuotas().getWithResponse("eastus", QuotaName.EXECUTION_MINUTES,
            com.azure.core.util.Context.NONE);
    }
}
