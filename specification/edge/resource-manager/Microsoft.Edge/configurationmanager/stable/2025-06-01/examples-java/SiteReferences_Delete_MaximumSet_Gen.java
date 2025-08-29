
/**
 * Samples for SiteReferences Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SiteReferences_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: SiteReferences_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void siteReferencesDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.siteReferences().delete("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
