
import com.azure.resourcemanager.workloadorchestration.models.SiteReferenceProperties;

/**
 * Samples for SiteReferences CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SiteReferences_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: SiteReferences_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void siteReferencesCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.siteReferences().define("testname").withExistingContext("rgconfigurationmanager", "testname")
            .withProperties(new SiteReferenceProperties().withSiteId("xxjpxdcaumewwgpbwzkcrgrcw")).create();
    }
}
