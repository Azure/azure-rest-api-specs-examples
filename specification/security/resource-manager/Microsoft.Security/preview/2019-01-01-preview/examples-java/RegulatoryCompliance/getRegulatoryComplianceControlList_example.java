
/**
 * Samples for RegulatoryComplianceControls List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/
     * RegulatoryCompliance/getRegulatoryComplianceControlList_example.json
     */
    /**
     * Sample code: Get all regulatory compliance controls details and state for selected standard.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getAllRegulatoryComplianceControlsDetailsAndStateForSelectedStandard(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.regulatoryComplianceControls().list("PCI-DSS-3.2", null, com.azure.core.util.Context.NONE);
    }
}
