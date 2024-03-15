
/**
 * Samples for AssessmentsMetadata List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/
     * ListAssessmentsMetadata_example.json
     */
    /**
     * Sample code: List security assessment metadata.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecurityAssessmentMetadata(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.assessmentsMetadatas().list(com.azure.core.util.Context.NONE);
    }
}
