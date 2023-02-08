import com.azure.resourcemanager.reservations.models.QuotaProperties;
import com.azure.resourcemanager.reservations.models.ResourceName;
import com.azure.resourcemanager.reservations.models.ResourceType;

/** Samples for Quota CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/putMachineLearningServicesQuotaRequestDedicated.json
     */
    /**
     * Sample code: Quotas_Request_PutForMachineLearningServices_DedicatedResource.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void quotasRequestPutForMachineLearningServicesDedicatedResource(
        com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .quotas()
            .define("StandardDv2Family")
            .withExistingLocation("D7EC67B3-7657-4966-BFFC-41EFD36BAAB3", "Microsoft.MachineLearningServices", "eastus")
            .withProperties(
                new QuotaProperties()
                    .withLimit(200)
                    .withUnit("Count")
                    .withName(new ResourceName().withValue("StandardDv2Family"))
                    .withResourceType(ResourceType.DEDICATED))
            .create();
    }
}
