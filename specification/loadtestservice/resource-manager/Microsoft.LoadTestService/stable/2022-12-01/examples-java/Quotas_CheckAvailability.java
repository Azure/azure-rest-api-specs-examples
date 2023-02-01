import com.azure.resourcemanager.loadtesting.models.QuotaBucketRequest;
import com.azure.resourcemanager.loadtesting.models.QuotaBucketRequestPropertiesDimensions;

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
    public static void quotasCheckAvailability(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
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
                com.azure.core.util.Context.NONE);
    }
}
