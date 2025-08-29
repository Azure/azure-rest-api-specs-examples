
/**
 * Samples for SiteReferences Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SiteReferences_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SiteReferences_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void siteReferencesGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.siteReferences().getWithResponse("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
