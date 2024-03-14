
/**
 * Samples for Assessments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/
     * ListAssessments_example.json
     */
    /**
     * Sample code: List security assessments.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecurityAssessments(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.assessments().list("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            com.azure.core.util.Context.NONE);
    }
}
