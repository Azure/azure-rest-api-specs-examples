
/**
 * Samples for PlaywrightQuotas ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/PlaywrightQuotas_ListBySubscription.json
     */
    /**
     * Sample code: PlaywrightQuotas_ListBySubscription.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void
        playwrightQuotasListBySubscription(com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.playwrightQuotas().listBySubscription("eastus", com.azure.core.util.Context.NONE);
    }
}
