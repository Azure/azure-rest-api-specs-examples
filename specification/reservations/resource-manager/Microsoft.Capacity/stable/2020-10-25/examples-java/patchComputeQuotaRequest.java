import com.azure.resourcemanager.reservations.models.CurrentQuotaLimitBase;
import com.azure.resourcemanager.reservations.models.QuotaProperties;
import com.azure.resourcemanager.reservations.models.ResourceName;

/** Samples for Quota Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/patchComputeQuotaRequest.json
     */
    /**
     * Sample code: Quotas_Request_PatchForCompute.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void quotasRequestPatchForCompute(
        com.azure.resourcemanager.reservations.ReservationsManager manager) {
        CurrentQuotaLimitBase resource =
            manager
                .quotas()
                .getWithResponse(
                    "D7EC67B3-7657-4966-BFFC-41EFD36BAAB3",
                    "Microsoft.Compute",
                    "eastus",
                    "standardFSv2Family",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new QuotaProperties()
                    .withLimit(200)
                    .withUnit("Count")
                    .withName(new ResourceName().withValue("standardFSv2Family")))
            .apply();
    }
}
