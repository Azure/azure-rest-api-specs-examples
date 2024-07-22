
/**
 * Samples for ResourceGuards List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * ResourceGuardCRUD/GetResourceGuardsInSubscription.json
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
