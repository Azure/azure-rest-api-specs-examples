/** Samples for Quota List. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getMachineLearningServicesQuotaLimits.json
     */
    /**
     * Sample code: Quotas_listQuotaLimitsMachineLearningServices.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasListQuotaLimitsMachineLearningServices(
        com.azure.resourcemanager.quota.QuotaManager manager) {
        manager
            .quotas()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus",
                com.azure.core.util.Context.NONE);
    }
}
