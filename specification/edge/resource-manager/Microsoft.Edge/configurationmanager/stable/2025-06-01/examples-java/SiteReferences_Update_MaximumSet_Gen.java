
import com.azure.resourcemanager.workloadorchestration.models.SiteReference;
import com.azure.resourcemanager.workloadorchestration.models.SiteReferenceProperties;

/**
 * Samples for SiteReferences Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SiteReferences_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: SiteReferences_Update_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void siteReferencesUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        SiteReference resource = manager.siteReferences()
            .getWithResponse("rgconfigurationmanager", "testname", "testname", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new SiteReferenceProperties().withSiteId("nwiuyaro")).apply();
    }
}
