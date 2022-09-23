import com.azure.core.util.Context;

/** Samples for AssessmentsMetadata DeleteInSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/DeleteAssessmentsMetadata_subscription_example.json
     */
    /**
     * Sample code: Delete a security assessment metadata for subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteASecurityAssessmentMetadataForSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .assessmentsMetadatas()
            .deleteInSubscriptionWithResponse("ca039e75-a276-4175-aebc-bcd41e4b14b7", Context.NONE);
    }
}
