import com.azure.resourcemanager.reservations.models.QuotaProperties;
import com.azure.resourcemanager.reservations.models.ResourceName;

/** Samples for Quota CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/putComputeOneSkuQuotaRequest.json
     */
    /**
     * Sample code: Quotas_Request_PutForCompute.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void quotasRequestPutForCompute(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .quotas()
            .define("standardFSv2Family")
            .withExistingLocation("D7EC67B3-7657-4966-BFFC-41EFD36BAAB3", "Microsoft.Compute", "eastus")
            .withProperties(
                new QuotaProperties()
                    .withLimit(200)
                    .withUnit("Count")
                    .withName(new ResourceName().withValue("standardFSv2Family")))
            .create();
    }
}
