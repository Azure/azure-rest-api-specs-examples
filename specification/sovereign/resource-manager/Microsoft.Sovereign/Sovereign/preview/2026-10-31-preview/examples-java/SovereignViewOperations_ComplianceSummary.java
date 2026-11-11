
/**
 * Samples for SovereignViewOperations ComplianceSummary.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-31-preview/SovereignViewOperations_ComplianceSummary.json
     */
    /**
     * Sample code: SovereignViewOperations_ComplianceSummary.
     * 
     * @param manager Entry point to SovereignManager.
     */
    public static void
        sovereignViewOperationsComplianceSummary(com.azure.resourcemanager.sovereign.SovereignManager manager) {
        manager.sovereignViewOperations().complianceSummaryWithResponse("ExampleResourceGroup", "ExampleSV",
            com.azure.core.util.Context.NONE);
    }
}
