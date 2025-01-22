
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/
     * getMachineLearningServicesUsages.json
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
