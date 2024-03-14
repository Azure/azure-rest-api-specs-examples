
/**
 * Samples for RegulatoryComplianceControls Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/
     * RegulatoryCompliance/getRegulatoryComplianceControl_example.json
     */
    /**
     * Sample code: Get selected regulatory compliance control details and state.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSelectedRegulatoryComplianceControlDetailsAndState(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.regulatoryComplianceControls().getWithResponse("PCI-DSS-3.2", "1.1", com.azure.core.util.Context.NONE);
    }
}
