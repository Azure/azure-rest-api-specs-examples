
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/getMachineLearningServicesUsages.json
     */
    /**
     * Sample code: Quotas_listUsagesMachineLearningServices.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasListUsagesMachineLearningServices(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.usages().list(
            "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus",
            com.azure.core.util.Context.NONE);
    }
}
