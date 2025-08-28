
import com.azure.resourcemanager.workloadorchestration.models.VersionParameter;

/**
 * Samples for ConfigTemplates RemoveVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ConfigTemplates_RemoveVersion_MaximumSet_Gen.json
     */
    /**
     * Sample code: ConfigTemplates_RemoveVersion_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void configTemplatesRemoveVersionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.configTemplates().removeVersionWithResponse("rgconfigurationmanager", "testname",
            new VersionParameter().withVersion("ghtvdzgmzncaifrnuumg"), com.azure.core.util.Context.NONE);
    }
}
