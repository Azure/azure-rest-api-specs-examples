import com.azure.resourcemanager.reservations.models.QuotaProperties;
import com.azure.resourcemanager.reservations.models.ResourceName;
import com.azure.resourcemanager.reservations.models.ResourceType;

/** Samples for Quota CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/putMachineLearningServicesQuotaRequestLowPriority.json
     */
    /**
     * Sample code: Quotas_Request_PutForMachineLearningServices_LowPriorityResource.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void quotasRequestPutForMachineLearningServicesLowPriorityResource(
        com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .quotas()
            .define("TotalLowPriorityCores")
            .withExistingLocation("D7EC67B3-7657-4966-BFFC-41EFD36BAAB3", "Microsoft.MachineLearningServices", "eastus")
            .withProperties(
                new QuotaProperties()
                    .withLimit(200)
                    .withUnit("Count")
                    .withName(new ResourceName().withValue("TotalLowPriorityCores"))
                    .withResourceType(ResourceType.LOW_PRIORITY))
            .create();
    }
}
