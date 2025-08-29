
/**
 * Samples for SiteReferences ListByContext.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SiteReferences_ListByContext_MaximumSet_Gen.json
     */
    /**
     * Sample code: SiteReferences_ListByContext_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void siteReferencesListByContextMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.siteReferences().listByContext("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
