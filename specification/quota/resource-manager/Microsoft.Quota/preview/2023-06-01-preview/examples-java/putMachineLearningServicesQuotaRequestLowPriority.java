
import com.azure.resourcemanager.quota.models.LimitObject;
import com.azure.resourcemanager.quota.models.QuotaProperties;
import com.azure.resourcemanager.quota.models.ResourceName;

/**
 * Samples for Quota CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/
     * putMachineLearningServicesQuotaRequestLowPriority.json
     */
    /**
     * Sample code: Quotas_Request_ForMachineLearningServices_LowPriorityResource.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasRequestForMachineLearningServicesLowPriorityResource(
        com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.quotas().define("TotalLowPriorityCores").withExistingScope(
            "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.MachineLearningServices/locations/eastus")
            .withProperties(new QuotaProperties().withLimit(new LimitObject().withValue(10))
                .withName(new ResourceName().withValue("TotalLowPriorityCores")).withResourceType("lowPriority"))
            .create();
    }
}
