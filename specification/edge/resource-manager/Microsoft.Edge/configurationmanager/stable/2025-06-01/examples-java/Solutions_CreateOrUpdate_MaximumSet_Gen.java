
import com.azure.resourcemanager.workloadorchestration.models.ExtendedLocation;
import com.azure.resourcemanager.workloadorchestration.models.ExtendedLocationType;
import com.azure.resourcemanager.workloadorchestration.models.SolutionProperties;

/**
 * Samples for Solutions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Solutions_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Solutions_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionsCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutions().define("testname").withExistingTarget("rgconfigurationmanager", "testname")
            .withProperties(new SolutionProperties()).withExtendedLocation(
                new ExtendedLocation().withName("szjrwimeqyiue").withType(ExtendedLocationType.EDGE_ZONE))
            .create();
    }
}
