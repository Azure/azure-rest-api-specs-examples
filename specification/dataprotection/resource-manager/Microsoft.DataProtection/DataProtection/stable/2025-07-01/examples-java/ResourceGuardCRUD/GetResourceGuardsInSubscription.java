
/**
 * Samples for ResourceGuards List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ResourceGuardCRUD/GetResourceGuardsInSubscription.json
     */
    /**
     * Sample code: Get ResourceGuards in Subscription.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getResourceGuardsInSubscription(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().list(com.azure.core.util.Context.NONE);
    }
}
