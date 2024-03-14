
/**
 * Samples for RegulatoryComplianceStandards Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/
     * RegulatoryCompliance/getRegulatoryComplianceStandard_example.json
     */
    /**
     * Sample code: Get selected regulatory compliance standard details and state.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSelectedRegulatoryComplianceStandardDetailsAndState(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.regulatoryComplianceStandards().getWithResponse("PCI-DSS-3.2", com.azure.core.util.Context.NONE);
    }
}
