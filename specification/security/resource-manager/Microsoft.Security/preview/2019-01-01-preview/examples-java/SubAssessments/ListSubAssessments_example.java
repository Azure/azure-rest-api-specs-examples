
/**
 * Samples for SubAssessments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/
     * ListSubAssessments_example.json
     */
    /**
     * Sample code: List security sub-assessments.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecuritySubAssessments(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.subAssessments().list("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            "82e20e14-edc5-4373-bfc4-f13121257c37", com.azure.core.util.Context.NONE);
    }
}
