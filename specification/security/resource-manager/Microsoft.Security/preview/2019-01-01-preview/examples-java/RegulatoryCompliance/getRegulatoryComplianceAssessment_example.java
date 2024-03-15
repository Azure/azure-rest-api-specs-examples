
/**
 * Samples for RegulatoryComplianceAssessments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/
     * RegulatoryCompliance/getRegulatoryComplianceAssessment_example.json
     */
    /**
     * Sample code: Get selected regulatory compliance assessment details and state.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSelectedRegulatoryComplianceAssessmentDetailsAndState(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.regulatoryComplianceAssessments().getWithResponse("PCI-DSS-3.2", "1.1",
            "968548cb-02b3-8cd2-11f8-0cf64ab1a347", com.azure.core.util.Context.NONE);
    }
}
