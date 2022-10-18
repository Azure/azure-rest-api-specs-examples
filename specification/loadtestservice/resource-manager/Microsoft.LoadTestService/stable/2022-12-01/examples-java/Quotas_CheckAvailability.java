import com.azure.core.util.Context;
import com.azure.resourcemanager.loadtestservice.models.QuotaBucketRequest;
import com.azure.resourcemanager.loadtestservice.models.QuotaBucketRequestPropertiesDimensions;

/** Samples for Quotas CheckAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_CheckAvailability.json
     */
    /**
     * Sample code: Quotas_CheckAvailability.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void quotasCheckAvailability(com.azure.resourcemanager.loadtestservice.LoadTestManager manager) {
        manager
            .quotas()
            .checkAvailabilityWithResponse(
                "westus",
                "testQuotaBucket",
                new QuotaBucketRequest()
                    .withCurrentUsage(20)
                    .withCurrentQuota(40)
                    .withNewQuota(50)
                    .withDimensions(
                        new QuotaBucketRequestPropertiesDimensions()
                            .withSubscriptionId("testsubscriptionId")
                            .withLocation("westus")),
                Context.NONE);
    }
}
