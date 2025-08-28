
import com.azure.resourcemanager.workloadorchestration.models.VersionParameter;

/**
 * Samples for Schemas RemoveVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Schemas_RemoveVersion_MaximumSet_Gen.json
     */
    /**
     * Sample code: Schemas_RemoveVersion_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemasRemoveVersionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemas().removeVersionWithResponse("rgconfigurationmanager", "testname",
            new VersionParameter().withVersion("ghtvdzgmzncaifrnuumg"), com.azure.core.util.Context.NONE);
    }
}
