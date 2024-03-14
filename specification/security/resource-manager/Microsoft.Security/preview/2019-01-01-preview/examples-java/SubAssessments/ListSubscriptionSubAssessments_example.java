
/**
 * Samples for SubAssessments ListAll.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/
     * ListSubscriptionSubAssessments_example.json
     */
    /**
     * Sample code: List security sub-assessments.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecuritySubAssessments(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.subAssessments().listAll("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            com.azure.core.util.Context.NONE);
    }
}
