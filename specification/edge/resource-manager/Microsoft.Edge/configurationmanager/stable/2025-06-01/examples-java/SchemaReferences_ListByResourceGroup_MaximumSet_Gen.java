
/**
 * Samples for SchemaReferences ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SchemaReferences_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: SchemaReferences_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemaReferencesListByResourceGroupMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemaReferences().listByResourceGroup("jdvtghygpz", com.azure.core.util.Context.NONE);
    }
}
