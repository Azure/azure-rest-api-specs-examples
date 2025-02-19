
/**
 * Samples for AssessmentsMetadata GetInSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/
     * GetAssessmentsMetadata_subscription_example.json
     */
    /**
     * Sample code: Get security assessment metadata for subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getSecurityAssessmentMetadataForSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.assessmentsMetadatas().getInSubscriptionWithResponse("21300918-b2e3-0346-785f-c77ff57d243b",
            com.azure.core.util.Context.NONE);
    }
}
