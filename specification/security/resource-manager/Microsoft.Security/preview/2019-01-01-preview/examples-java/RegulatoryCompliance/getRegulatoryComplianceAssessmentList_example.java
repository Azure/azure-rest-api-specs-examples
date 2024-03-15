
/**
 * Samples for RegulatoryComplianceAssessments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/
     * RegulatoryCompliance/getRegulatoryComplianceAssessmentList_example.json
     */
    /**
     * Sample code: Get all assessments mapped to selected regulatory compliance control.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getAllAssessmentsMappedToSelectedRegulatoryComplianceControl(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.regulatoryComplianceAssessments().list("PCI-DSS-3.2", "1.1", null, com.azure.core.util.Context.NONE);
    }
}
